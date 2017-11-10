package server

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/masakistewart/GoBotty/handlers"
	"github.com/masakistewart/GoBotty/router"
	. "github.com/masakistewart/GoBotty/logging"
)

func StartServer() {
	// backingService holds all the variables needed in other parts of the app
	// backingService := setup()

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	r := router.New(handlers.Root)

	r.Handle("GET", "/", handlers.Root)
	r.Handle("GET", "/users", handlers.Users)
	r.Handle("GET", "/users/:id", handlers.User)
	http.ListenAndServe(":8080", r)

	Trace.Println("Started the server on Port 8080")
}
