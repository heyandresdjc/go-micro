package main

type notification struct {
	Sender  string `json:"sender"`
	Reciver string `json:"reciver"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Notifications []notification

var notifications = Notifications{}

func RepoCreateNotifcation(noti notification) notification {
	notifications = append(notifications, noti)
	return noti
}
