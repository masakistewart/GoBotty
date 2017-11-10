package handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

func User(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "hello %s", params["id"])
}
