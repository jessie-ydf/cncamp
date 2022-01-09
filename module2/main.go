package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/debug/pprof", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal("start server failed, error: %s \n", err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to cloud native</h1>"))

	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)

	for name, headers := range r.Header {
		for _, header := range headers {
			fmt.Printf("Header key: %s, Header value: %s \n", name, header)
			w.Header().Set(name, header)
		}
	}

	clientip := getCurrentIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")

	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}
