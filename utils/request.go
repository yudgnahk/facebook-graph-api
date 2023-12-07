package utils

import "net/http"

func AddJsonHeader(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
}

func AddHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Add(key, value)
	}
}
