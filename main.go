package main

import (
	"log"
	"net/http"
	"time"
)

var config *Config

func main() {

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
	response, code, cacheExists, endpoint := validateCache(r.Method, r.URL)
	if cacheExists {
		w.WriteHeader(code)
		w.Write([]byte(response))
		return
	}
	requestURI := r.URL.RequestURI()
	response, code, headers := accessOriginalPath(r, requestURI)
	if endpoint != nil {
		writeCache(endpoint, requestURI, response, code, headers)
	}
	for name, value := range headers {
		w.Header().Set(name, value[0])
	}
	w.WriteHeader(code)
	w.Write([]byte(response))
}
