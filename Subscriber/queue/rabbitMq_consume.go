package queue

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const CreateQueue = "publisher.create"

type Consume struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewConsumeEvent(conn *amqp.Connection, ch *amqp.Channel) *Consume {
	return &Consume{conn: conn, ch: ch}
}

func (c *Consume) PrintEvent() {			//currently listening only to create queue
	q, err := c.ch.QueueDeclare(
		CreateQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		failOnError(err)
	}

	msgs, err := c.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err)
	}

	isAlive := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-isAlive
}

func failOnError(err error) {
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}
