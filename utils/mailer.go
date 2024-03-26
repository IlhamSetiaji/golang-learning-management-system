package utils

import (
	"encoding/json"
	"log"
	"net"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailerService struct {
	HostPort string
	User     string
	Passcode string
}
type message struct {
	To      string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

func (ms *MailerService) SendMail(jsonBody []byte) {
	var msg message

	if err := json.Unmarshal(jsonBody, &msg); err != nil {
		log.Fatal(err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", ms.User)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)

	m.SetBody("text/html", msg.Body)
	host, port_str, _ := net.SplitHostPort(ms.HostPort)
	port_number, _ := strconv.Atoi(port_str)
	d := gomail.NewDialer(host, port_number, ms.User, ms.Passcode)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
