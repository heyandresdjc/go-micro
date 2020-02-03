package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAllNotification(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notifications)
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var newNotification notification

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &newNotification); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if len(os.Getenv("SENDGRID_API_KEY")) > 0 {
		if err := SendGridEmail(
			newNotification.Sender,
			newNotification.Reciver,
			newNotification.Subject,
			newNotification.Body); err != nil {
			panic(err)
		}
	} else {
		if err := SendRelayEmail(
			newNotification.Sender,
			newNotification.Reciver,
			newNotification.Subject,
			newNotification.Body); err != nil {
			panic(err)
		}
	}

	if err := json.NewEncoder(w).Encode(newNotification); err != nil {
		panic(err)
	}
}
