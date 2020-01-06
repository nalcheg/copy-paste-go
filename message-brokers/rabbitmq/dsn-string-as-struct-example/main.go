package main

import (
	"log"

	"github.com/brianvoe/gofakeit"

	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq"
	"github.com/streadway/amqp"
)

type DSN string

func (dsn DSN) GetValueOfDsnStringStruct() string {
	return string(dsn)
}

func (dsn *DSN) GetPointToDsnStringStruct() string {
	return string(*dsn)
}

func (dsn DSN) ConnectRabbitmq() (*amqp.Connection, *amqp.Channel, error) {
	gofakeit.Seed(0)

	// methods are equal, just an example
	if gofakeit.Number(0, 100) > 50 {
		return ConnectRabbitmq(dsn.GetValueOfDsnStringStruct())
	} else {
		return ConnectRabbitmq(dsn.GetPointToDsnStringStruct())
	}
}

func ConnectRabbitmq(dsn string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

func main() {
	var dsn DSN
	dsn = rabbitmq.Dsn

	conn, ch, err := dsn.ConnectRabbitmq()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := ch.Close(); err != nil {
			log.Fatal(err)
		}
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := rabbitmq.PrepareRabbimq(ch); err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		rabbitmq.QueueName,
		"",
		false,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "dead-letters",
			"x-dead-letter-routing-key": "dead-letters.routing.key",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for m := range msgs {
		log.Print(string(m.Body))
		if err := m.Ack(true); err != nil {
			log.Fatal(err)
		}
	}
	log.Fatalf(" [*] RabbitMQ consume error, we are terminating")
}
