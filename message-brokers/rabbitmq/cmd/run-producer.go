package main

import (
	"log"

	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq/connector"
	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq/producer"
)

func main() {
	rabbitmqConn, err := connector.RabbitmqConnector{}.ConnectRabbitmq(dsn)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	prod := func(rabbitmqConn *connector.RabbitmqConnector) producer.Interface {
		produrer := producer.NewProducer(rabbitmqConn)

		return produrer
	}(rabbitmqConn)

	if err := prod.CreateRouting(exchangeName, routingKey, queueName); err != nil {
		log.Fatalf("%#v", err)
	}

	if err := prod.Send(exchangeName, routingKey, []byte("mess")); err != nil {
		log.Fatalf("%#v", err)
	}
}
