package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IlhamSetiaji/go-lms/internal/config"
	"github.com/IlhamSetiaji/go-lms/internal/model"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Consumer struct {
	Channel     *amqp091.Channel
	Queue       string
	Log         *logrus.Logger
	MailService *config.MailService
}

func (c *Consumer) Consume(ctx context.Context) {
	msgs, err := c.Channel.Consume(
		c.Queue,
		"",
		false, // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.Log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var emailEvent model.EmailEvent
			if err := json.Unmarshal(d.Body, &emailEvent); err != nil {
				c.Log.Printf("Error decoding email event: %v", err)
				continue
			}

			data := config.MailData{
				From:    emailEvent.From,
				To:      []string{emailEvent.To},
				Subject: emailEvent.Subject,
				Body:    emailEvent.Body,
			}

			if err := c.MailService.SendMail(data); err != nil {
				c.Log.Printf("Error sending email: %v", err)
				// If an error occurred, reject the message and requeue it
				if err := d.Nack(false, true); err != nil {
					c.Log.Printf("Error nacking message: %v", err)
				}
				continue
			}

			// If the message was successfully processed, acknowledge it
			if err := d.Ack(false); err != nil {
				c.Log.Printf("Error acknowledging message: %v", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	viper := viper.New()
	logger := logrus.New()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()

	if err != nil {
		logger.Fatalf("Fatal error config file: %v", err)
	}

	conn, err := amqp091.Dial(viper.GetString("rabbitmq.url"))
	if err != nil {
		logger.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		logger.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	mailService := config.NewMailService(logger, viper)

	consumer := &Consumer{
		Channel:     channel,
		Queue:       "users",
		Log:         logger,
		MailService: mailService,
	}

	consumer.Consume(context.Background())
}
