package util

import (
	"gopkg.in/gomail.v2"
)

type GomailConf struct {
	Smtp     string `yaml:"smtp"`
	Port     int    `yaml:"port"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

type GoMailDialer struct {
	*gomail.Dialer
	Username string
}

func NewGoMail(conf *GomailConf) *GoMailDialer {
	dialer := &GoMailDialer{
		Dialer:   gomail.NewDialer(conf.Smtp, conf.Port, conf.Account, conf.Password),
		Username: conf.Account,
	}
	return dialer
}

func (gmd *GoMailDialer) DialAndSend(messages []*gomail.Message) error {
	return gmd.Dialer.DialAndSend(messages...)
}

type MailMessage struct {
	From    string          `json:"from" validate:"required"`
	To      []string        `json:"to" validate:"required"`
	Cc      []string        `json:"cc"`
	Subject string          `json:"subject" validate:"required"`
	Body    MailMessageBody `json:"body"`
}

type MailMessageBody struct {
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
}

func (mm *MailMessage) NewMessage() *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", mm.From)
	m.SetHeader("To", mm.To...)
	if len(mm.Cc) > 0 {
		m.SetHeader("Cc", mm.Cc...)
	}
	m.SetHeader("Subject", mm.Subject)
	m.SetBody(mm.Body.ContentType, mm.Body.Content)
	return m
}
