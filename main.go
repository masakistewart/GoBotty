package main

import (
	"log"
	"net/http"

	"github.com/masakistewart/GoBotty/handlers"
)

func main() {
	_port := ":3030"
	log.Printf("Server has started on localhost%s", _port)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/results", handlers.ResultsHandler)
	http.HandleFunc("/test", handlers.Test)
	http.ListenAndServe(_port, nil)
}
