package messaging

import (
	"github.com/IlhamSetiaji/go-lms/internal/model"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type UserProducer struct {
	Producer[*model.UserEvent]
}

func NewUserProducer(channel *amqp091.Channel, log *logrus.Logger) *UserProducer {
	return &UserProducer{
		Producer: Producer[*model.UserEvent]{
			Channel: channel,
			Queue:   "users",
			Log:     log,
		},
	}
}
