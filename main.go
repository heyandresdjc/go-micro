package main

import (
	"log"
	"net/http"
)

func main() {
	Config()
	Routes := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", Routes))
}
