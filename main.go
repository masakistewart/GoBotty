package main

import (
	"net/http"

<<<<<<< HEAD
	"github.com/masakistewart/GoBotty/discordApi"
=======
	"./handlers"
>>>>>>> 88dddb4875f36e59bb4f5480fe0c4eee57d42e4e
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/results", handlers.ResultsHandler)
	http.ListenAndServe(":3030", nil)
}
