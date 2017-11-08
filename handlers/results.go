package handlers

import (
	"encoding/json"
	"net/http"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting response to json", http.StatusTeapot)
	}
	w.Write(jsonBody)
}
