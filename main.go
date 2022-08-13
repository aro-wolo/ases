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
	To       string
	Subject  string
}

func Sendmail(h Header, msg string) {

	m := gomail.NewMessage()
	m.SetHeader("From", h.From)
	m.SetHeader("To", h.To)
	m.SetHeader("Subject", h.Subject)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(h.Host, 587, h.Username, h.Password)
	d.DialAndSend(m)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

}

/* func main() {

	h := Header{}
	h.Username = "AKIAIE6URQG7YVNUOSIA"
	h.Password = "AlH8zkhRZh3rs+G3wDa4pWkJ3NOzjHi7a5tos9Hpr8nl"
	h.Host = "email-smtp.us-west-2.amazonaws.com"
	h.From = "do-not-reply@officelime.app"
	h.Subject = "Testing"
	h.To = "adexfe@live.com"

	Sendmail(h, "Hello world")

	println("hi...")

}
*/
