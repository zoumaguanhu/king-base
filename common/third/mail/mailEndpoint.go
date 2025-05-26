package mail

import (
	"gopkg.in/gomail.v2"
)

type MailConf struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Username        string `json:"userName"`
	Password        string `json:"password"`
	SSL             bool   `json:"ssl"`
	WorkerNum       int    `json:"workerNum"`
	JobQueue        int    `json:"jobQueue"`
	WaitTimeOut     int    `json:"waitTimeOut"`
	RetryCount      int    `json:"retryCount"`
	MaxCountPerDay  int    `json:"maxCountPerDay"`
	MaxCountPerFile int    `json:"maxCountPerFile"`
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
func (s *MailEndpoint) GetConfig() *MailConf {
	return s.mailConf
}
func (s *MailEndpoint) GetDialer() *gomail.Dialer {
	return s.dialer
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
