package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/masakistewart/GoBotty/discordApi"
)

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println(os.Getenv("DISCORD_TOKEN"))
	discordApi.GetMe()
}
