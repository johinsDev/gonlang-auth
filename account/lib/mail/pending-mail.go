package mail

type PendingMail struct {
	mailer *Mailer
	to     struct {
		Name    string
		Address string
	}
}

func (p *PendingMail) To(address string, name ...string) *PendingMail {
	p.to = struct {
		Name    string
		Address string
	}{
		Name:    name[0],
		Address: address,
	}

	return p
}

func (p *PendingMail) Send(mailable MailableContract) {
	p.mailer.SendMailable(mailable, p.to)
}
