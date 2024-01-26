package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	//create a simple http server to respond to hello route
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal("Error on listening to PORT 8080")
	}

}

func hello(w http.ResponseWriter, r *http.Request) {

	//Use of anonymous struct
	message := struct {
		Status string `json:"status"`
		Body   string `json:"body"`
	}{
		Status: "success",
		Body:   "We got your request!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
