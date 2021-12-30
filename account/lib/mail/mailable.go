package mail

import "fmt"

type MailableContract interface {
	Send(mailer *Mailer, message *Message) string
	Build(m *Message) *Message
}

type Mailable struct{}

func (m *Mailable) Send(mailer *Mailer, message *Message) string {
	fmt.Println(message)
	return "a"
}
