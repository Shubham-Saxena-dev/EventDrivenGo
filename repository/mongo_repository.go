package repository

import (
	"GoEvents/queue"
	"GoEvents/requests"
	"context"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAllEmployees() ([]requests.AccountCreateRequest, error)
	GetAccount() (requests.AccountCreateRequest, error)
	CreateAccount() error
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
	GET_QUEUE = "publisher.get"
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

	q := createQueues(GET_QUEUE, r.conn, r.ch)

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

func (r repository) GetAccount() (requests.AccountCreateRequest, error) {
	panic("implement me")
}

func (r repository) CreateAccount() error {
	panic("implement me")
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
