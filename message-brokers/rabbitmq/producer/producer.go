package producer

import (
	"github.com/nalcheg/copy-paste-go/message-brokers/rabbitmq/connector"
	"github.com/streadway/amqp"
)

type Producer struct {
	conn  *amqp.Connection
	chann *amqp.Channel
}

func NewProducer(c *connector.RabbitmqConnector) *Producer {
	return &Producer{conn: c.Conn, chann: c.Chan}
}

type Interface interface {
	Send(exchName, rKey string, message []byte) error
	CreateRouting(exchName, rKey, queueName string) error
}

func (p *Producer) Send(exchName, rKey string, message []byte) error {
	return p.chann.Publish(
		exchName,
		rKey,
		false,
		false,
		amqp.Publishing{
			Body:         message,
			DeliveryMode: 2,
		},
	)
}

func (p *Producer) CreateRouting(exchName, rKey, queueName string) error {
	if err := p.chann.ExchangeDeclare(
		exchName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	queue, err := p.chann.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return p.chann.QueueBind(
		queue.Name,
		rKey,
		exchName,
		false,
		nil,
	)
}
