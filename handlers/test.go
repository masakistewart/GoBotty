package handlers

import (
	"encoding/json"
	"github.com/masakistewart/GoBotty/logging"
	"net/http"
	"net/url"
)

type FormData map[string]string

func Test(w http.ResponseWriter, r *http.Request, params url.Values) {
	r.ParseForm()
	logging.Info.Println(r.Form)
	t := new(FormData)
	for key, _ := range r.Form {
        logging.Info.Println(key)
        //LOG: {"test": "that"}
        err := json.Unmarshal([]byte(key), &t)
        if err != nil {
            logging.Error.Println(err.Error())
        }
    }
}
