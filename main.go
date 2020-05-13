package main

import (
	"log"
	"net/http"
	"quiz/api"
)

// create tests
// create api documentation
func main() {
	r := api.NewRouter()
	log.Fatal("Server is working on port: 5000", http.ListenAndServe(":5000", r))
}
