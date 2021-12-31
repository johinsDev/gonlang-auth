package mails

import (
	"github.com/johinsDev/authentication/lib/mail"
	"github.com/johinsDev/authentication/models"
)

type Welcome struct {
	User *models.User
	mail.Mailable
}

func (m *Welcome) Build() {
	m.Subject("Testing from mailable").
		From("noreply@johinsdev.com").
		Layout([]string{"layout"}).
		View([]string{"layout.html", "template.html"}).
		With(struct {
			Name string
			URL  string
		}{
			Name: m.User.Name,
			URL:  "testing.com",
		})
}
