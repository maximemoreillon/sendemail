package main

import (
    "fmt"
		"os"
    "strconv"

    gomail "gopkg.in/mail.v2"
)

func main() {

	smtpHost:= os.Getenv("SMTP_HOST")
	smtpUsername:= os.Getenv("SMTP_USERNAME")
	smtpPassword:= os.Getenv("SMTP_PASSWORD")
	smtpFrom:= os.Getenv("SMTP_FROM")
	smtpTo:= os.Getenv("SMTP_TO")
	smtpSubject:= os.Getenv("SMTP_SUBJECT")
	smtpMessage:= os.Getenv("SMTP_MESSAGE")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		panic(err)
	}

	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", smtpFrom)
	message.SetHeader("To", smtpTo)
	message.SetHeader("Subject", smtpSubject)
	message.SetBody("text/plain", smtpMessage)

	// Set up the SMTP dialer
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
			fmt.Println("Error:", err)
			panic(err)
	} else {
			fmt.Println("Email sent successfully!")
	}
}