package main

import (
	"log"
	"net/http"
	"quiz/api"
)

func main() {
	r := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))

}
