package main

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/nalcheg/copy-paste-go/self-made-apigateway/consts"
	"github.com/nalcheg/copy-paste-go/self-made-apigateway/rabbitmq"
)

func main() {
	rabbitConn, rabbitCh, err := rabbitmq.Connect(consts.RabbitmqDSN)
	if err != nil {
		log.Panic(err)
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	if err := rabbitmq.Prepare(rabbitCh); err != nil {
		log.Panic(err)
	}

	msgs, err := rabbitCh.Consume(
		consts.RequestsQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	for m := range msgs {
		if err := rabbitCh.Publish(
			consts.ExchangeName,
			consts.ResponsesRoutingKey,
			false,
			false,
			amqp.Publishing{
				ContentType:  "application/json",
				Body:         m.Body,
				DeliveryMode: 2,
				Headers:      m.Headers,
			},
		); err != nil {
			log.Panic(err)
		}
		if err := m.Ack(true); err != nil {
			log.Panic(err)
		}
	}
}
