package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	fmt.Println("reqpayload",requestPayload)
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("readjson hatası")
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//validate the user againist the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		fmt.Println("validate gmail hatası")
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		fmt.Println("validate password hatası")
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
