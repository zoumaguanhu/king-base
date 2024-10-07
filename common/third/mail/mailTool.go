package mail

import "crypto/tls"

type MailConf struct {
	MailSmtp  string
	MailPort  int64
	SysEmail  string
	SysPasswd string
	SslVerify bool
}
type MailTarget struct {
	mailConf MailConf
	ToList   []string
	MsgBody  string
	Subject  string
	Meta     string
}

func (l *MailTarget) SendMail(subject, body string, to []string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "your_email@example.com")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.example.com", 587, "your_email@example.com", "your_email_password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}
