package messaging

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
	Log     *log.Logger
}

func NewProducer(amqpURL string, queueName string) (*Producer, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Producer{Channel: ch, Queue: q}, nil
}

func (p *Producer) Publish(body string) error {
	ctx := context.Background() // Or use a different context if you have one
	err := p.Channel.PublishWithContext(
		ctx,
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		p.Log.Fatal(err)
		return err
	}
	return nil
}
