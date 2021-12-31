package mail

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"strconv"

	"github.com/johinsDev/authentication/config"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Auth smtp.Auth
	From struct {
		Name    string
		Address string
	}
	Config *config.MaiLConfig
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

func (m *Mailer) To(address string, name ...string) *PendingMail {
	pedingMail := &PendingMail{
		mailer: m,
		to: struct {
			Name    string
			Address string
		}{},
	}

	pedingMail.To(address, name...)

	return pedingMail
}

func (m *Mailer) SendMailable(
	mailable MailableContract,
	to struct {
		Name    string
		Address string
	},
) {

	mailable.Build()

	mailable.To(to.Address).Send(m)
}

func (m *Mailer) Send(
	views []string,
	data interface{},
	cb ...func(message *Message),
) {
	message := m.buildMessage()

	callback := cb[0]

	buf, _ := m.parseView(views, data, message.layout)

	message.Body(buf.String())

	if callback != nil {
		callback(message)
	}

	// SMTP DRIVER
	PORT, err := strconv.Atoi(m.Config.PORT)

	if err != nil {
		log.Fatal("Error loading PORT")
	}

	d := gomail.NewDialer(
		m.Config.HOST,
		PORT,
		m.Config.USERNAME,
		m.Config.PASSWORD,
	)

	s, err := d.Dial()

	messageGoMail := gomail.NewMessage()

	messageGoMail.SetHeader("From", message.from.Address, message.from.Name)

	for _, to := range message.to {
		messageGoMail.SetAddressHeader("To", to.Address, to.Name)
	}

	messageGoMail.SetHeader("Subject", message.subject)

	messageGoMail.SetBody("text/html", message.body)

	for _, file := range message.attachments {
		messageGoMail.Attach(file)
	}

	if err := gomail.Send(s, messageGoMail); err != nil {
		log.Fatal("Error sending mail", err)
	}
}

func (m Mailer) parseView(
	views []string,
	data interface{},
	layout []string,
) (*bytes.Buffer, *template.Template) {

	t, err := template.ParseFiles(views...)

	if err != nil {
		log.Fatal("Error sending mail", err)
	}

	var buf bytes.Buffer

	t.ExecuteTemplate(&buf, "layout", data)

	for _, v := range layout {
		t.ExecuteTemplate(&buf, v, data)
	}

	return &buf, t
}

func (m Mailer) buildMessage() *Message {
	message := &Message{}

	if m.From.Address != "" {
		message.From(m.From.Address, m.From.Name)
	}

	return message
}

// Constructor

func NewMailer() *Mailer {
	config := config.GetMailConfig()

	auth := smtp.PlainAuth(
		"",
		config.USERNAME,
		config.PASSWORD,
		config.HOST,
	)

	mailer := &Mailer{
		Auth:   auth,
		Config: config,
	}

	return mailer.Alwaysfrom(config.FROM.ADDRESS, config.FROM.NAME)
}
