package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Create a Simple REST API and implement rate limiter
func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:8080", nil)

	time.Sleep(500 * time.Millisecond)

}

func hello(w http.ResponseWriter, r *http.Request) {
	// Create anonymous struct(we need to use it onle once)
	message := struct {
		Message string `json:"message"`
		Body    string `json:"body"`
	}{
		Message: "success",
		Body:    "A Simple Rest API",
	}
	w.Header().Set("content-Type", "applicatio/json")
	w.WriteHeader(http.StatusOK)

	//check if request is allowed or not
	ra := rate.Every(2 * time.Second)

	limiter := rate.NewLimiter(ra, 1)
	if !limiter.Allow() {
		message.Message = "failed"
		message.Body = "Too Many request"
		w.WriteHeader(http.StatusTooManyRequests)

	}
	json.NewEncoder(w).Encode(message)
	//time.Sleep(200 * time.Millisecond)
}
