package connector

import "github.com/streadway/amqp"

type RabbitmqConnector struct {
	Conn *amqp.Connection
	Chan *amqp.Channel
}

func (rc RabbitmqConnector) ConnectRabbitmq(dsn string) (*RabbitmqConnector, error) {
	var err error

	rc.Conn, err = amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	rc.Chan, err = rc.Conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rc, nil
}
