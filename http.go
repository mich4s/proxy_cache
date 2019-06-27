package main

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

func prepareResponse(w http.ResponseWriter, r *http.Request) (string, int, http.Header) {
	cached, cacheExists, endpoint := validateCache(r.Method, r.URL)
	if cacheExists {
		return cached.response, cached.code, cached.headers
	}
	requestURI := r.URL.RequestURI()
	response, code, headers := accessOriginalPath(r, requestURI)
	if endpoint != nil {
		writeCache(endpoint, requestURI, response, code, headers)
	}
	return response, code, headers
}

func accessOriginalPath(r *http.Request, requestURI string) (string, int, http.Header) {
	client := &http.Client{}
	log.Println(config.PrivateURL + requestURI)
	request, err := http.NewRequest(r.Method, config.PrivateURL+requestURI, r.Body)
	if err != nil {
		return "", 500, nil
	}
	fillRequestHeaders(request, r.Header)
	resp, err := client.Do(request)
	if err != nil {
		return "", 500, nil
	}
	responseString := convertResponseToString(resp)
	return responseString, resp.StatusCode, resp.Header
}

func fillRequestHeaders(request *http.Request, headers http.Header) {
	for name, headers := range headers {
		name = strings.ToLower(name)
		for _, h := range headers {
			request.Header.Set(name, h)
		}
	}
}

func convertResponseToString(response *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String()
}
