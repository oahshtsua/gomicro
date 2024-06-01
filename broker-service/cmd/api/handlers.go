package main

import (
	"log"
	"net/http"

	js "github.com/oahshtsua/gomicro/broker-service/pkg/json"
)

func brokerHandler(w http.ResponseWriter, r *http.Request) {
	resp := js.Response{
		Error:   false,
		Message: "Hit the broker!",
	}

	err := js.Write(w, http.StatusAccepted, resp, nil)
	if err != nil {
		log.Fatal(err)
	}
}
