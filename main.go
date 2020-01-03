package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type Variables struct {
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to this life-changing API.")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)

	setUp()
	SendTextEmail("ajc2101@gmail.com", "heyandresdjc@gmail.com", "Really?!", "this ia a test msg")
}

func setUp() {
	jsonFile, err := os.Open("./vars/env.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var variables Variables
	json.Unmarshal(byteValue, &variables)
	for i := 0; i < len(variables.Variables); i++ {
		os.Setenv(variables.Variables[i].Name, variables.Variables[i].Value)
	}
	fmt.Println("Env Setup")
}

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
