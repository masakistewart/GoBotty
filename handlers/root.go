package handlers

import "net/http"
import "net/url"
import "github.com/masakistewart/GoBotty/logging"
import "fmt"

func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	w.WriteHeader(200)
	ip := r.Referer()
	logging.Info.Printf("The root route was hit by %s", ip)
	logging.Info.Println(r)
	fmt.Fprintf(w,"IP: %s", ip)
}
