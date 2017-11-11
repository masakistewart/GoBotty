package handlers

import (
	"github.com/masakistewart/GoBotty/logging"
	"net/http"
	"net/url"
)

func Test(w http.ResponseWriter, r *http.Request, params url.Values) {
	r.ParseForm()
	logging.Info.Println(r.Form)
	defer r.Body.Close()
}
