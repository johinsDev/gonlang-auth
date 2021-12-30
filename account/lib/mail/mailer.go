package mail

import (
	"bytes"
	"fmt"
	"html/template"
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

func (m *Mailer) getAddr() string {
	return fmt.Sprintf("%s:%s", m.Config.HOST, m.Config.PORT)
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
	message := &Message{}

	message.SetTo(to.Address, to.Name)

	mailable.Send(m, mailable.Build(message))
}

func (m *Mailer) Send(
	views []string,
	data interface{},
	cb ...func(message *Message, template *template.Template),
) {
	message := m.buildMessage()

	callback := cb[0]

	buf, t := m.parseView(views, data)

	if callback != nil {
		callback(message, t)
	}

	message.SetBody(buf.String())

	err := smtp.SendMail(
		m.getAddr(),
		m.Auth,
		message.From,
		message.To,
		message.GetBody(),
	)

	if err != nil {
		log.Fatal("Error sending mail", err)
	}

	fmt.Println("message.To", message.To)
}

func (m Mailer) parseView(views []string, data interface{}) (*bytes.Buffer, *template.Template) {

	t, err := template.ParseFiles(views...)

	if err != nil {
		log.Fatal("Error sending mail", err)
	}

	buf := new(bytes.Buffer)

	t.ExecuteTemplate(buf, "layout", nil)

	if err = t.Execute(buf, data); err != nil {
		log.Fatal("Error sending mail", err)
	}

	return buf, t
}

func (m Mailer) buildMessage() *Message {
	message := &Message{}
	return message.SetFrom(m.From.Address, m.From.Name)
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
