package server

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/masakistewart/GoBotty/handlers"
	"github.com/masakistewart/GoBotty/router"
	"github.com/masakistewart/GoBotty/logging"
)

// StartServer starts the server and initializes the rest of the app
func StartServer() {
	// backingService holds all the variables needed in other parts of the app
	// backingService := setup()
	PORT := ":8080"
	logging.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	r := router.New(handlers.Root)

	r.Handle("GET", "/", handlers.Root)
	r.Handle("GET", "/users", handlers.Users)
	r.Handle("GET", "/users/:id", handlers.User)
	r.Handle("GET", "/test", handlers.Test)
	
	logging.Info.Printf("Started the server on Port %s", PORT)
	http.ListenAndServe(PORT, r)
}
