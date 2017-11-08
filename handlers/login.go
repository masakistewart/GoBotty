package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/masakistewart/GoBotty/discordApi"
	"github.com/masakistewart/GoBotty/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred models.Credentials
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
	w.WriteHeader(http.StatusOK)
}
