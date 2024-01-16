package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Endpoint string
	Port     string
	Exchange string
	Socket   *amqp.Connection
	Channel  *amqp.Channel
	Queue    amqp.Queue
}
