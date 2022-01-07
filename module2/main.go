package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", headers)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, header := range headers {
			w.Header().Add(name, header)
		}
	}
}
