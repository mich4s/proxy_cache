package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func validateCache(method string, url *url.URL) (*Cached, bool, *Endpoint) {
	endpoint := findEndpoint(url.RequestURI())
	if endpoint == nil {
		return nil, false, endpoint
	}
	cache, exists := findCacheElement(endpoint, url.RequestURI())
	if !exists {
		return nil, false, endpoint
	}
	return cache, true, endpoint
}

func findCacheElement(endpoint *Endpoint, url string) (*Cached, bool) {
	validateCacheWasInitialized(endpoint)
	for _, cache := range endpoint.cached {
		if cache.URL == url {
			return cache, true
		}
	}
	return nil, false
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
	headers = prepareHeaders(headers)
	endpoint.cached = append(endpoint.cached, &Cached{
		URL:        requestURI,
		response:   response,
		code:       code,
		lastUpdate: time.Now().Unix(),
		headers:    headers,
	})
}

func prepareHeaders(headers http.Header) http.Header {
	for name, headers := range headers {
		name = strings.ToLower(name)
		for _, h := range headers {
			fmt.Println(name, h)
		}
	}
	return headers
}
