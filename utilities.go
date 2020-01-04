package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendTextEmail(sender string, reciver string, subject string, body string) {
	email_user := os.Getenv("EMAIL_USER")
	email_pass := os.Getenv("EMAIL_PASS")
	email_host := os.Getenv("EMAIL_HOST")
	addr := email_host + ":587"
	auth := smtp.PlainAuth("", email_user, email_pass, email_host)
	err := smtp.SendMail(addr, auth, sender, []string{reciver}, []byte(body))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent")
}
