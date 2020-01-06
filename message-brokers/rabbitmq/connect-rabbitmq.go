package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	Dsn          = "amqp://guest:guest@127.0.0.1:5672"
	ExchangeName = "exchange.name"
	RoutingKey   = "routing.key"
	QueueName    = "queue.name"
)

func PrepareRabbimq(ch *amqp.Channel) error {
	if err := ch.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	if err := ch.QueueBind(
		queue.Name,
		RoutingKey,
		ExchangeName,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
