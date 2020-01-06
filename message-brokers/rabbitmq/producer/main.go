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
		if err := ch.Close(); err != nil {
			log.Fatal(err)
		}
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := ch.Publish(
		rabbitmq.ExchangeName,
		rabbitmq.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte("first txt message"),
			DeliveryMode: 2,
		},
	); err != nil {
		log.Fatal(err)
	}

	if err := ch.Publish(
		rabbitmq.ExchangeName,
		rabbitmq.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte("second txt message"),
			DeliveryMode: 2,
		},
	); err != nil {
		log.Fatal(err)
	}
}
