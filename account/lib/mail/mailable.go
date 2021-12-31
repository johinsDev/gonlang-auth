package mail

import (
	"html/template"
)

type MailableContract interface {
	Send(mailer *Mailer)
	Build()
	To(address string) *Mailable
}

type Mailable struct {
	to      []string
	from    string
	subject string
	view    []string
	layout  []string
	data    interface{}
}

func (m *Mailable) To(address string) *Mailable {
	m.to = append(m.to, address)
	return m
}

func (m *Mailable) Subject(subject string) *Mailable {
	m.subject = subject
	return m
}

func (m *Mailable) From(address string) *Mailable {
	m.from = address
	return m
}

func (m *Mailable) View(view []string) *Mailable {

	m.view = append(m.view, view...)
	return m
}

func (m *Mailable) Layout(layout []string) *Mailable {
	m.layout = append(m.layout, layout...)
	return m
}

func (m *Mailable) With(data interface{}) *Mailable {
	m.data = data
	return m
}

func (m *Mailable) Send(mailer *Mailer) {
	mailer.Send(m.view, m.data, func(message *Message, template *template.Template) {
		message.Subject(m.subject)

		for _, to := range m.to {
			message.To(to)
		}

		if m.from != "" {
			message.From(m.from)
		}
	})
}
