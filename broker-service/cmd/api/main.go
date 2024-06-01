package main

import (
	"log"
	"net/http"
)

func main() {

	srv := &http.Server{
		Addr:    ":5000",
		Handler: routes(),
	}

	log.Println("Starting server on port: 5000")
	log.Fatal(srv.ListenAndServe())
}
