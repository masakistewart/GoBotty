package server

import (
	"fmt"
	"os"
	"strconv"
)

// UserCredentials are for login/cookie info
type UserCredentials struct {
	Email    string
	Password string
}

// EndPoints are the endpoint for an API
type EndPoints struct {
	BaseURL string
	UsersEP string
}

// Client holds API credentials
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
