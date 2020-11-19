package main

import (
	"encoding/json"
	"flexmessage"
	"log"
	"net/http"
)

const addr = "localhost:8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var notify flexmessage.FlexMessage

		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			notify.Error(r.Method + " method is not implemented")
		}
		if notify.NoErrors() {
			notify.Message("Ok")

		}
		json.NewEncoder(w).Encode(notify.Compact())

	})

	log.Fatal(http.ListenAndServe(addr, nil))

}
