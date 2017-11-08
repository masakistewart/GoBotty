package handlers

import (
	"encoding/json"
	"net/http"
)

type placeHolderError struct {
	Message string
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	err := placeHolderError{Message: "Sorry for the inconveniance but this route is currently down :("}
	w.Header()
	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(err)
	return
}
