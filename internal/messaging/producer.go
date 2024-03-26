package messaging

import (
	"context"
	"encoding/json"

	"github.com/IlhamSetiaji/go-lms/internal/model"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Producer[T model.Event] struct {
	Channel *amqp091.Channel
	Queue   string
	Log     *logrus.Logger
}

func (p *Producer[T]) GetQueue() *string {
	return &p.Queue
}

func (p *Producer[T]) Send(ctx context.Context, event T) error {
	// Declare the queue before publishing messages
	_, err := p.Channel.QueueDeclare(
		"users", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		p.Log.WithError(err).Error("failed to declare queue")
		return err
	}

	bodyBytes, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	err = p.Channel.PublishWithContext(
		ctx,
		"",      // exchange
		p.Queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
			Headers: amqp091.Table{
				"id": event.GetEmail(),
			},
		})
	if err != nil {
		p.Log.WithError(err).Error("failed to publish message")
		return err
	}

	return nil
}
