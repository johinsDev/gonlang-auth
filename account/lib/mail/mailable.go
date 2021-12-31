package mail

type MailableContract interface {
	Send(mailer *Mailer)
	Build()
	To(address string, name ...string) *Mailable
}

type Mailable struct {
	to []struct {
		Name    string
		Address string
	}
	from struct {
		Name    string
		Address string
	}
	subject     string
	view        []string
	layout      []string
	data        interface{}
	attachments []string
}

func (m *Mailable) GetName(name ...string) string {
	Name := ""

	if len(name) >= 1 {
		Name = name[0]
	}

	return Name
}

func (m *Mailable) To(address string, name ...string) *Mailable {

	m.to = append(m.to, struct {
		Name    string
		Address string
	}{
		Name:    m.GetName(name...),
		Address: address,
	})

	return m
}

func (m *Mailable) Subject(subject string) *Mailable {
	m.subject = subject
	return m
}

func (m *Mailable) From(address string, name ...string) *Mailable {

	m.from = struct {
		Name    string
		Address string
	}{
		Name:    m.GetName(name...),
		Address: address,
	}

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

func (m *Mailable) Attach(file string) *Mailable {
	m.attachments = append(m.attachments, file)

	return m
}

func (m *Mailable) Send(mailer *Mailer) {
	mailer.Send(m.view, m.data, func(message *Message) {
		message.Subject(m.subject)

		for _, to := range m.to {
			message.To(to.Address, to.Name)
		}

		for _, file := range m.attachments {
			message.Attach(file)
		}

		if m.from.Address != "" {
			message.From(m.from.Address, m.from.Name)
		}

		message.Layout(m.layout)
	})
}
