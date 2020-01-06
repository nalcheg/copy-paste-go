package main

import (
	"log"

	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	conn, ch, err := rabbitmq.ConnectRabbitmq(rabbitmq.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
		if err := ch.Close(); err != nil {
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
	}

	log.Fatalf(" [*] RabbitMQ consume error, we are terminating")
}
