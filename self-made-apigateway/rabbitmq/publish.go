package rabbitmq

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(ch *amqp.Channel) *Publisher {
	return &Publisher{ch: ch}
}

func (p *Publisher) Publish(exch, rKey string, id uuid.UUID, msg []byte) error {
	headers := make(map[string]interface{})
	headers["reqID"] = id.String()

	return p.ch.Publish(
		exch,
		rKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         msg,
			DeliveryMode: 2,
			Headers:      headers,
		},
	)
}
