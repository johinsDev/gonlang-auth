package mails

import (
	"github.com/johinsDev/authentication/lib/mail"
	"github.com/johinsDev/authentication/models"
)

type Welcome struct {
	User *models.User
	mail.Mailable
}

// Todo support ses
// manager
// switch driver
// queue or later
func (m *Welcome) Build() {
	m.
		CC("otyheremail@gmail.com", "testing@gmail.com").
		Subject("Welcome to codecourse").
		From("noreply@codecourse.com", "No reply").
		Layout([]string{"layout"}).
		View([]string{"template.html", "layout.html"}).
		Attach("models/user.go").
		With(struct {
			Name string
			URL  string
		}{
			Name: m.User.Name,
			URL:  "testing.com",
		})
}
