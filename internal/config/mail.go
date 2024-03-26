package config

import (
	"crypto/tls"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type MailService struct {
	Log     *logrus.Logger
	Viper   *viper.Viper
	To      string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

func NewMailService(log *logrus.Logger, viper *viper.Viper) *MailService {
	return &MailService{
		Log:   log,
		Viper: viper,
	}
}

type MailData struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Body    string
	Attach  string
}

func (ms *MailService) SendMail(data MailData) error {
	m := gomail.NewMessage()

	m.SetHeader("From", data.From)

	m.SetHeader("To", data.To...)

	for _, cc := range data.Cc {
		m.SetAddressHeader("Cc", cc, "")
	}

	m.SetHeader("Subject", data.Subject)

	m.SetBody("text/html", data.Body)

	if data.Attach != "" {
		m.Attach(data.Attach)
	}

	d := gomail.NewDialer(ms.Viper.GetString("mail.host"), ms.Viper.GetInt("mail.port"), ms.Viper.GetString("mail.username"), ms.Viper.GetString("mail.password"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
