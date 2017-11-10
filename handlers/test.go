package handlers

import (
	"github.com/masakistewart/GoBotty/logging"
	"net/http"
	"net/url"
)

func Test(w http.ResponseWriter, r *http.Request, params url.Values) {
	if 1 != 2 {
		logging.Error.Println("1 is not 2 all is right with the world")
		w.WriteHeader(404)
	}
}
