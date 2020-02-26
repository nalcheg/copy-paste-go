package rabbitmq

import (
	"github.com/streadway/amqp"

	"github.com/nalcheg/copy-paste-go/self-made-apigateway/consts"
)

func Prepare(ch *amqp.Channel) error {
	if err := ch.ExchangeDeclare(
		consts.ExchangeName,
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
		consts.RequestsQueue,
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
		consts.RequestsRoutingKey,
		consts.ExchangeName,
		false,
		nil,
	); err != nil {
		return err
	}

	queue, err = ch.QueueDeclare(
		consts.ResponsesQueue,
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
		consts.ResponsesRoutingKey,
		consts.ExchangeName,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
