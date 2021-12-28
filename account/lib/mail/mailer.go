package mail

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/johinsDev/authentication/config"
)

type Mailer struct {
	Auth smtp.Auth
	From struct {
		Name    string
		Address string
	}
}

func (m *Mailer) Alwaysfrom(address string, name ...string) *Mailer {
	m.From = struct {
		Name    string
		Address string
	}{
		Name:    name[0],
		Address: address,
	}

	return m
}

func (m *Mailer) Send(view string, data interface{}, cb ...func(message *Message)) {
	message := m.buildMessage()

	callback := cb[0]

	if callback != nil {
		callback(message)
	}

	config := config.GetMailConfig()

	msg := []byte("To: " + message.To[0] + "\r\n" +
		"Subject: Why are you not using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", config.HOST, config.PORT),
		m.Auth,
		message.From,
		message.To,
		msg,
	)

	if err != nil {
		log.Fatal("Error sending mail", err)
	}

	fmt.Println("message.To", message.To)
}

func (m Mailer) buildMessage() *Message {
	message := &Message{}
	return message.SetFrom(m.From.Address, m.From.Name)
}

func NewMailer() *Mailer {
	config := config.GetMailConfig()

	auth := smtp.PlainAuth(
		"",
		config.USERNAME,
		config.PASSWORD,
		config.HOST,
	)

	mailer := &Mailer{
		Auth: auth,
	}

	return mailer.Alwaysfrom(config.FROM.ADDRESS, config.FROM.NAME)
}
