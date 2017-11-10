package handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

func Users(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "not done yet")
}
