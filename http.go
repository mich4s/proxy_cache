package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func accessOriginalPath(r *http.Request, requestURI string) (string, int, http.Header) {
	fmt.Println("Writing api request")
	client := &http.Client{}
	request, err := http.NewRequest("GET", config.PrivateURL+requestURI, nil)
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
			request.Header.Set(validateHeader(name, h))
		}
	}
}

func validateHeader(name string, value string) (string, string) {
	if name == "accept-encoding" {
		value = strings.ReplaceAll(value, "gzip, ", "")
		value = strings.ReplaceAll(value, "gzip", "")
	}
	fmt.Println(name, value)
	return name, value
}

func omitGZip() {

}

func convertResponseToString(response *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String()
}
