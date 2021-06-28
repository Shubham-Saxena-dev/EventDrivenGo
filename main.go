package main

import (
	"GoEvents/controllers"
	"GoEvents/repository"
	"GoEvents/routes"
	"GoEvents/service"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Database   = "goMongos"
	Collection = "accounts"
	MongoDbUrl = "mongodb://mongodb:27017/"
)

var (
	collection *mongo.Collection
	repo       repository.Repository
	serv       service.Service
	controller controllers.Controller
	ctx        context.Context
	conn       *amqp.Connection
	ch         *amqp.Channel
)

func main() {
	log.Info("Hi, this is event publishing and consuming")
	configQueue()
	initDatabase()
	createServer()
}

func createServer() {

	server := gin.Default()
	initializeLayers()

	routes.RegisterHandlers(server, controller).RegisterHandlers()

	err := server.Run()

	if err != nil {
		failOnError(err, "Unable to start server")
	}
}

func initializeLayers() {
	repo = repository.NewMongoRepository(collection, ctx, conn, ch)
	serv = service.NewRepository(repo)
	controller = controllers.NewController(serv)
}

func configQueue() {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to rabbit mq")
	ch, err = conn.Channel()
	failOnError(err, "Failed to create channel")
}

func initDatabase() {
	log.Info("Connecting to datastore")
	clientOptions := options.Client().ApplyURI(MongoDbUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		failOnError(err, "Unable to establish connection with MongoDB")
	}

	collection = client.Database(Database).Collection(Collection)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
		panic(err)
	}
}
