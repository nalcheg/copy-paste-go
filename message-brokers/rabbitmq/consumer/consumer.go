package consumer

import (
	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq/connector"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn  *amqp.Connection
	chann *amqp.Channel
}

func NewConsumer(c *connector.RabbitmqConnector) *Consumer {
	return &Consumer{conn: c.Conn, chann: c.Chan}
}

type Interface interface {
	Consume(queueName string, ch chan []byte) error
	CreateQueue(queueName string) error
}

func (c *Consumer) Consume(queueName string, ch chan []byte) error {
	msgs, err := c.chann.Consume(
		queueName,
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
		return err
	}

	for m := range msgs {
		ch <- m.Body
		if err := m.Ack(true); err != nil {
			return err
		}
	}

	return nil
}

func (c *Consumer) CreateQueue(queueName string) error {
	_, err := c.chann.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}
