package mail

import "fmt"

type Message struct {
	From    string
	To      []string
	Subject string
}

func (m *Message) SetFrom(adddress string, name ...string) *Message {
	m.From = fmt.Sprintf("%s", adddress)

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
