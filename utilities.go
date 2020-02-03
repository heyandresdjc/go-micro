package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Config() {
	env_file := "vars/.env"
	data, err := ioutil.ReadFile(env_file)
	if err != nil {
		log.Fatal(err)
	}

	key_pair_value_list := strings.Split(string(data), "\n")
	for i := 0; i < len(key_pair_value_list); i++ {
		variable := strings.Split(key_pair_value_list[i], "=")
		os.Setenv(string(variable[0]), variable[1])
	}
}

func SendRelayEmail(sender string, reciver string, subject string, body string) error {

	email_user := os.Getenv("EMAIL_HOST_USER")
	email_pass := os.Getenv("EMAIL_HOST_PASSWORD")
	email_host := os.Getenv("EMAIL_HOST")
	email_port := os.Getenv("EMAIL_PORT")
	addr := email_host + ":" + email_port

	auth := smtp.PlainAuth("", email_user, email_pass, email_host)

	err := smtp.SendMail(addr, auth, sender, []string{reciver}, []byte(body))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent")
	return err
}

func SendGridEmail(sender string, reciver string, subject string, body string) error {

	from := mail.NewEmail("Example User", sender)
	to := mail.NewEmail("Example User", reciver)
	plainTextContent := ""
	htmlContent := body
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return err
}
