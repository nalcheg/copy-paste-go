package rabbitmq

import "github.com/streadway/amqp"

func Connect(dsn string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}
