package ases

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Header struct {
	Username string
	From     string
	Password string
	Host     string
}

func Sendmail(to string, sub string, msg string, h Header) {

	m := gomail.NewMessage()
	m.SetHeader("From", h.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(h.Host, 587, h.Username, h.Password)
	d.DialAndSend(m)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

}
