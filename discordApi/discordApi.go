package discordApi

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	// Email user email for discord
	Email string
	// Password for discord
	Password string
	// Token for discord authentication
	Token string
)

// GetToken gets token from discord
func GetToken(credentials struct {
	Email    string
	Password string
}) {
	dg, err := discordgo.New(credentials.Email, credentials.Password)
	if err != nil {
		fmt.Println("Error while creating discord session:", err)
		return
	}

	Token := dg.Token

	fmt.Printf("Your Authentication Token is:\n\n%s\n", Token)
}

// CheckEnvs for environment (env) variables
func CheckEnvs() error {
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

func checkToken() bool {
	if os.Getenv("DISCORDTOKEN") != "" {
		log.Println("No Token Found")
		return false
	}
	return true
}

func getMe() {
	http.Get("https://discordapp.com/api/v6/users/s1ler")
}
