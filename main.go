package main

import (
	"log"
	"net/http"
	"time"
)

var config *Config

func main() {
	New()
}

func New() {

	loadConfig()

	loadCacheRunner()

	server := &http.Server{
		Addr:         ":" + config.PublicPort,
		WriteTimeout: config.WriteTimeout * time.Second,
		ReadTimeout:  config.ReadTimeout * time.Second,
		Handler:      http.HandlerFunc(handleRequest),
	}
	log.Fatal(server.ListenAndServe())
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	response, code, headers := prepareResponse(w, r)
	writeResponse(w, response, code, headers)
}

func writeResponse(w http.ResponseWriter, response string, code int, headers http.Header) {
	for name, value := range headers {
		w.Header().Set(name, value[0])
	}
	w.WriteHeader(code)
	w.Write([]byte(response))
}
