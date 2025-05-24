package mail

import (
	"gopkg.in/gomail.v2"
)

type MailConf struct {
	Host     string
	Port     int
	Username string
	Password string
	SSL      bool
}
type MailEndpoint struct {
	mailConf *MailConf
	dialer   *gomail.Dialer
}

func New(c *MailConf) *MailEndpoint {
	dialer := gomail.NewDialer(
		c.Host,
		c.Port,
		c.Username,
		c.Password,
	)
	dialer.SSL = c.SSL

	return &MailEndpoint{mailConf: c, dialer: dialer}
}

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
	IsHTML  bool
}

func (s *MailEndpoint) SendEmail(email *Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To...)
	m.SetHeader("Subject", email.Subject)

	if email.IsHTML {
		m.SetBody("text/html", email.Body)
	} else {
		m.SetBody("text/plain", email.Body)
	}
	return s.dialer.DialAndSend(m)
}
