package mail

type Message struct {
	from struct {
		Name    string
		Address string
	}
	to struct {
		Name    string
		Address string
	}
	subject     string
	body        string
	view        []string
	layout      []string
	attachments []string
	cc          []string
}

func (m *Message) GetName(name ...string) string {
	Name := ""

	if len(name) >= 1 {
		Name = name[0]
	}

	return Name
}

func (m *Message) CC(address ...string) *Message {
	m.cc = append(m.cc, address...)

	return m
}

func (m *Message) Attach(file string) *Message {
	m.attachments = append(m.attachments, file)

	return m
}

func (m *Message) Layout(layout []string) *Message {
	m.layout = append(m.layout, layout...)

	return m
}

func (m *Message) From(adddress string, name ...string) *Message {
	m.from = struct {
		Name    string
		Address string
	}{
		Name:    m.GetName(name...),
		Address: adddress,
	}

	return m
}

func (m *Message) To(adddress string, name ...string) *Message {

	m.to = struct {
		Name    string
		Address string
	}{
		Name:    m.GetName(name...),
		Address: adddress,
	}

	return m
}

func (m *Message) Subject(subject string) *Message {
	m.subject = subject

	return m
}

func (m *Message) Body(body string) *Message {

	m.body = body

	return m
}

func (m *Message) GetBody() []byte {
	return []byte(m.body)
}
