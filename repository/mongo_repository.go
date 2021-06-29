package repository

import (
	"GoEvents/queue"
	"GoEvents/requests"
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAllEmployees() ([]requests.AccountCreateRequest, error)
	GetAccount(id string) (requests.AccountCreateRequest, error)
	CreateAccount(requests.AccountCreateRequest) (requests.AccountCreateRequest, error)
	UpdateAccount() error
	DeleteAccount() error
}

type repository struct {
	collection *mongo.Collection
	ctx        context.Context
	conn       *amqp.Connection
	ch         *amqp.Channel
}

const (
	GetQueue    = "publisher.get"
	CreateQueue = "publisher.create"
)

func NewMongoRepository(collection *mongo.Collection, ctx context.Context,
	conn *amqp.Connection, ch *amqp.Channel) Repository {

	return &repository{
		collection: collection,
		ctx:        ctx,
		conn:       conn,
		ch:         ch,
	}
}

func (r repository) GetAllEmployees() ([]requests.AccountCreateRequest, error) {

	q := createQueues(GetQueue, r.conn, r.ch)

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Requesting to get all accounts"),
	}

	publishMessage(r.ch, q.Name, msg)

	cursor, err := r.collection.Find(r.ctx, bson.D{})
	defer cursor.Close(r.ctx)

	if err != nil {
		return []requests.AccountCreateRequest{}, err
	}

	var accounts []requests.AccountCreateRequest

	if cursor.All(r.ctx, &accounts); err != nil {
		return []requests.AccountCreateRequest{}, err
	}

	return accounts, nil
}

func (r repository) GetAccount(id string) (requests.AccountCreateRequest, error) {
	q := createQueues(GetQueue, r.conn, r.ch)

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Requesting account for ID " + id),
	}

	publishMessage(r.ch, q.Name, msg)

	account := requests.AccountCreateRequest{}
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account, err
	}
	err = r.collection.FindOne(r.ctx, bson.M{
		"_id": uid,
	}).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (r repository) CreateAccount(request requests.AccountCreateRequest) (requests.AccountCreateRequest, error) {

	request.Id = primitive.NewObjectID()
	request.Dept.DeptId = primitive.NewObjectID()

	q := createQueues(CreateQueue, r.conn, r.ch)
	body, err := json.Marshal(request)
	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	publishMessage(r.ch, q.Name, msg)

	_, err = r.collection.InsertOne(r.ctx, request)

	if err != nil {
		return requests.AccountCreateRequest{}, err
	}
	return requests.AccountCreateRequest{
		Id: request.Id,
	}, nil
}

func (r repository) UpdateAccount() error {
	panic("implement me")
}

func (r repository) DeleteAccount() error {
	panic("implement me")
}

func createQueues(name string, conn *amqp.Connection, ch *amqp.Channel) amqp.Queue {
	return queue.NewQueue(name, conn, ch).CreateQueue()
}

//Default exchange
func publishMessage(ch *amqp.Channel, name string, msg amqp.Publishing) {
	ch.Publish("", name, false, false, msg)
}
