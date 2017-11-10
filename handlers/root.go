package handlers

import "net/http"
import "fmt"
import "net/url"

func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "This is Root")
}
