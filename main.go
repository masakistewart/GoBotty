package main

import (
	"errors"
	"fmt"
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

func main() {
	err := checkEnvs()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
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
