package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllNotification(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notifications)
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var newNotification notification
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(reqBody))
	json.Unmarshal(reqBody, &newNotification)
	notifications = append(notifications, newNotification)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newNotification)

}
