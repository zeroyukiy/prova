package mailer

import (
	"fmt"
	"net/smtp"
)

type Mailer struct {
	Msg string
}

func NewMailer() *Mailer {
	return &Mailer{
		Msg: "Messaggio di prova",
	}
}

func (e *Mailer) SendEmail(email string) {
	// Sender data.
	from := "513f514a-3911-4b62-b150-46d61555e640"
	password := "f671d502-2fd7-4d8f-b211-ebd1b2cde409"

	// Receiver email address.
	// to := []string{
	// 	"sender@example.com",
	// }

	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "app.debugmail.io"
	smtpPort := "25"

	// Message.
	message := []byte(e.Msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	go send(smtpHost, smtpPort, auth, from, to, message)
}

func send(host string, port string, auth smtp.Auth, from string, to []string, message []byte) {
	err := smtp.SendMail(host+":"+port, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}