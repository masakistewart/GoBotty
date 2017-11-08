package handlers

import (
	"encoding/json"
	"net/http"

	"../discordApi"
)

type Credentials struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	if r.Body == nil {
		http.Error(w, "Please send something in the json, Im lonely :(", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil || cred.Password == "" || cred.Email == "" {
		http.Error(w, "Please send something in the json, Im lonely :(", 400)
		return
	}

	discordApi.GetToken(cred)
}
