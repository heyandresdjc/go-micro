package main

type notification struct {
	Sender  string `json:"sender"`
	Reciver string `json:"reciver"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Notifications []notification

var notifications = Notifications{
	{
		Sender:  "a@a.com",
		Reciver: "a@a.com",
		Subject: "Test Subject",
		Body:    "Test Body",
	},
}
