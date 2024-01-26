package rabbitmq

import (
	"context"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbit() *Rabbit {
	return &Rabbit{
		Exchange: "",
	}
}

func (r *Rabbit) Configure() {

	r.Endpoint = "amqp://" + os.Getenv("RABBIT_USERNAME") + ":" + os.Getenv("RABBIT_PASSWORD") + "@" + os.Getenv("RABBIT_URL") + ":"
	r.Port = os.Getenv("RABBIT_AMQP_PROTOCOL")
}

func (r *Rabbit) ConnectSocket() {
	socket, err := amqp.Dial(r.Endpoint + r.Port)
	Must(err, "Failed to connect to RabbitMQ")
	log.Println("Socket connected")
	r.Socket = socket
	// defer conn.Close()
}

func (r *Rabbit) ConnectChannel() {

	channel, err := r.Socket.Channel()
	Must(err, "Failed to open a channel")
	log.Println("Channel Connected")
	r.Channel = channel
	// defer channel.Close()

}

func (r *Rabbit) DeclareQueue(queuename string) {

	queue, err := r.Channel.QueueDeclare(
		queuename, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	Must(err, "Failed to declare a queue")
	log.Println("QueueDeclared")
	r.Queue = queue
}

func (r *Rabbit) PublishCTXByte(ctx context.Context, data []byte) {

	err := r.Channel.PublishWithContext(
		ctx,
		r.Exchange,
		r.Queue.Name,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	Must(err, "Failed to publish the message")
	log.Printf(" [x] Sent %s\n", string(data))

}

func (r *Rabbit) Close() {
	err := r.Channel.Close()
	Must(err, "Cannot close channel")

	err = r.Socket.Close()
	Must(err, "Cannot close connection/socket")
}
