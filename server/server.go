package server

import (
	"net/http"

	"github.com/masakistewart/GoBotty/handlers"
	"github.com/masakistewart/GoBotty/router"
)

func StartServer() {
	// backingService holds all the variables needed in other parts of the app
	// backingService := setup()

	r := router.New(handlers.Root)

	r.Handle("GET", "/", handlers.Root)
	r.Handle("GET", "/users", handlers.Users)
	r.Handle("GET", "/users/:id", handlers.User)
	http.ListenAndServe(":8080", r)
}
