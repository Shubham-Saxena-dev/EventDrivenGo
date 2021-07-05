package main

import (
	"GoEvents/Subscriber/queue"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const AMQP_Url = "amqp://guest@localhost:5672"

var (
	conn *amqp.Connection
	ch   *amqp.Channel
)

func main() {
	log.Info("Hi, this is event consuming")
	configMessagingQueue()
	q := queue.NewConsumeEvent(conn, ch)
	q.PrintEvent()
}

func configMessagingQueue() {
	var err error
	conn, err = amqp.Dial(AMQP_Url)
	failOnError(err, "Failed to connect to rabbit mq")
	ch, err = conn.Channel()
	failOnError(err, "Failed to create channel")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
		panic(err)
	}
}
