package main

import (
	"log"

	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq/connector"
	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq/consumer"
)

func main() {
	rabbitmqConn, err := connector.RabbitmqConnector{}.ConnectRabbitmq(dsn)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	cons := func(rabbitmqConn *connector.RabbitmqConnector) consumer.Interface {
		cons := consumer.NewConsumer(rabbitmqConn)
		if err := cons.CreateQueue(queueName); err != nil {
			log.Fatalf("%#v", err)
		}

		return cons
	}(rabbitmqConn)

	ch := make(chan []byte, 10)

	go func() {
		for {
			select {
			case msg := <-ch:
				log.Printf("%s", msg)
			}
		}
	}()

	if err := cons.Consume(queueName, ch); err != nil {
		log.Fatalf("%#v", err)
	}
}
