package config

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewRabbitMQConsumer(config *viper.Viper, log *logrus.Logger) (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(config.GetString("rabbitmq.url"))
	if err != nil {
		log.Fatalf("Failed to create new connection: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, err
	}

	return ch, nil
}

func NewRabbitMQProducer(config *viper.Viper, log *logrus.Logger) (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(config.GetString("rabbitmq.url"))
	if err != nil {
		log.Fatalf("Failed to create new connection: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, err
	}

	return ch, nil
}
