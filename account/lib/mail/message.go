package mail

type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func (m *Message) SetFrom(adddress string, name ...string) *Message {
	m.From = adddress

	return m
}

func (m *Message) SetTo(adddress string, name ...string) *Message {
	m.To = append(m.To, adddress)

	return m
}

func (m *Message) SetSubject(subject string) *Message {
	m.Subject = subject

	return m
}

func (m *Message) SetBody(body string) *Message {
	subject := "Subject: " + m.Subject + "!\n"

	m.Body = subject + MIME + "\n" + body

	return m
}

func (m *Message) GetBody() []byte {
	return []byte(m.Body)
}
