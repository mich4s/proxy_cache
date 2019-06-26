package main

import (
	"net/http"
	"net/url"
	"time"
)

func validateCache(method string, url *url.URL) (string, int, bool, *Endpoint) {
	endpoint := findEndpoint(url.RequestURI())
	if endpoint == nil {
		return "", 0, false, endpoint
	}
	response, code, exists := findCacheElement(endpoint, url.RequestURI())
	if !exists {
		return "", 200, false, endpoint
	}
	return response, code, true, endpoint
}

func findCacheElement(endpoint *Endpoint, url string) (string, int, bool) {
	validateCacheWasInitialized(endpoint)
	for _, cache := range endpoint.cached {
		if cache.URL == url {
			return cache.response, cache.code, true
		}
	}
	return "", 0, false
}

func findEndpoint(url string) *Endpoint {
	for _, endpoint := range config.Enpoints {
		endpointLength := len(endpoint.URL)
		if endpoint.URL == url {
			return endpoint
		}
		if endpoint.URL == url[:endpointLength] {
			return endpoint
		}
	}
	return nil
}

func writeCache(endpoint *Endpoint, requestURI string, response string, code int, headers http.Header) {
	validateCacheWasInitialized(endpoint)
	insertCacheValue(endpoint, requestURI, response, code, headers)
}

func validateCacheWasInitialized(endpoint *Endpoint) {
	if len(endpoint.cached) == 0 {
		endpoint.cached = make([]*Cached, 0)
	}
}

func insertCacheValue(endpoint *Endpoint, requestURI string, response string, code int, headers http.Header) {
	endpoint.cached = append(endpoint.cached, &Cached{
		URL:        requestURI,
		response:   response,
		code:       code,
		lastUpdate: time.Now().Unix(),
		headers:    headers,
	})
}
