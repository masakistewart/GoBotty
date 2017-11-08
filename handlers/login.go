package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type credentials struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred credentials
	if r.Body == nil {
		http.Error(w, "Please send something in the json, Im lonely :(", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println(cred)
}
