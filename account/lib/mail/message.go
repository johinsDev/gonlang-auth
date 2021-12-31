package mail

type Message struct {
	from    string
	to      []string
	subject string
	body    string
	view    []string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func (m *Message) From(adddress string, name ...string) *Message {
	m.from = adddress

	return m
}

func (m *Message) To(adddress string, name ...string) *Message {
	m.to = append(m.to, adddress)

	return m
}

func (m *Message) Subject(subject string) *Message {
	m.subject = subject

	return m
}

func (m *Message) Body(body string) *Message {
	subject := "Subject: " + m.subject + "!\n"

	m.body = subject + MIME + "\n" + body

	return m
}

func (m *Message) GetBody() []byte {
	return []byte(m.body)
}
