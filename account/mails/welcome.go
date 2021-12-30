package mails

import (
	"github.com/johinsDev/authentication/lib/mail"
	"github.com/johinsDev/authentication/models"
)

type Welcome struct {
	user models.User
	mail.Mailable
}

func (m *Welcome) Build(message *mail.Message) *mail.Message {
	// review alwaysFrom
	// setView
	// and with method
	// setBuject, setTo,setFrom, setView, setWith wrap on mailable not use message, setLayout
	return message.SetFrom("johinsdev@gmail.com", "johan").SetSubject("Subject my dog")
	// return m.SetSubject("Subject").
	// 	SetView([]string{"template.html"}).
	// 	With(struct{ user models.User }{user: m.user})
}
