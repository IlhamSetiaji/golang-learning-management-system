package messaging

import (
	"github.com/IlhamSetiaji/go-lms/internal/model"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type EmailProducer struct {
	Producer[*model.EmailEvent]
}

func NewEmailProducer(channel *amqp091.Channel, log *logrus.Logger) *EmailProducer {
	return &EmailProducer{
		Producer: Producer[*model.EmailEvent]{
			Channel: channel,
			Queue:   "users",
			Log:     log,
		},
	}
}
