package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Variables for login
var (
	Email    string
	Password string
	Token    string
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

// checks for environment (env) variables
func checkEnvs() error {
	Email = os.Getenv("DISCORDEMAIL")
	Password = os.Getenv("DISCORDPASS")

	if Email == "" || Password == "" {
		return errors.New("Please add your discord email and password")
	}

	// the os can set an env
	os.Setenv("DISCORDTOKEN", Token)
	fmt.Println("Checking login information")
	return nil
}

func getToken() {

	dg, err := discordgo.New(Email, Password)
	if err != nil {
		fmt.Println("Error while creating discord session:", err)
		return
	}

	Token := dg.Token

	fmt.Printf("Your Authentication Token is:\n\n%s\n", Token)
}
