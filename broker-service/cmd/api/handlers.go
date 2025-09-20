package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		//Error: false,
		Message: "Hit the broker",
	}

	_ = app.writeSJON(w, http.StatusOK, payload)
}
