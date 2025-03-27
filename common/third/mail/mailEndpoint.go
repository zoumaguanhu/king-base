package mail

import (
	"gopkg.in/gomail.v2"
	"log"
)

type MailConf struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}
type MailEndpoint struct {
	mailConf MailConf
}

func (l *MailEndpoint) SendMail(to, subject, body string) error {
	cfg := l.mailConf
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
