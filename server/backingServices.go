package server

import (
	"fmt"
	"os"
	"strconv"
)

type UserCredentials struct {
	Email    string
	Password string
}

type EndPoints struct {
	BaseURL string
	UsersEP string
}

type Client struct {
	ID        int
	Secret    string
	UserLogin UserCredentials
	EndPoints EndPoints
}

// Setup all environment variables
func setup() Client {
	var (
		client          Client
		endPoints       EndPoints
		userCredentials UserCredentials
		email           string
		id, secret      string
	)

	email = os.Getenv("DISCORD_EMAIL")
	id = os.Getenv("DISCORD_CLIENT_ID")
	secret = os.Getenv("DISCORD_SECRET")

	client = Client{}
	client.ID, _ = strconv.Atoi(id)
	client.Secret = secret

	endPoints = EndPoints{}
	endPoints.BaseURL = os.Getenv("DISCORD_API_URL")
	endPoints.UsersEP = "/users"

	userCredentials = UserCredentials{}
	userCredentials.Email = email

	client.EndPoints = endPoints
	client.UserLogin = userCredentials

	fmt.Println(client)
	return client
}
