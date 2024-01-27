package main

import (
	"encoding/json"
	"net/http"
)

// Create a Simple REST API and implement rate limiter
func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:8080", nil)

}

func hello(w http.ResponseWriter, r *http.Request) {

	// Create anonymous struct(Creating here as we need to use onle once)

	message := struct {
		Message string `json:"message"`
		Body    string `json:"body"`
	}{
		Message: "success",
		Body:    "A Simple Rest API",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "applicatio/json")
	json.NewEncoder(w).Encode(message)

}
