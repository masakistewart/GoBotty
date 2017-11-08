package handlers

import (
	"net/http"

	"../discordApi"
)

func Test(w http.ResponseWriter, r *http.Request) {
	discordApi.GetMe()
}
