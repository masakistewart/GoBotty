package main

import (
	"net/http"

	"./handlers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/results", handlers.ResultsHandler)
	http.HandleFunc("/test", handlers.Test)
	http.ListenAndServe(":3030", nil)
}
