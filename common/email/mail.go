package email

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type Email struct {
	host 	 string
	port     int
	username string
	password string
	to       []string
	subj     string
	msg      string
}

func NewEmail(host, username, password string, port int) *Email {
	return &Email{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (e *Email) WithInfo(subj, msg string, to []string) {
	e.subj = subj
	e.msg = msg
	e.to = to
}

func (e *Email) Send() (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", e.username);
	m.SetHeader("To", e.to...);
	m.SetHeader("Subject", e.subj)
	m.SetBody("text/html", e.msg);
	d := gomail.NewDialer(e.host, e.port, e.username, e.password);
	if err = d.DialAndSend(m); err != nil {
		log.Error("send mail err: %+v. Method: SendMail#email", err);
	}
	return
}
