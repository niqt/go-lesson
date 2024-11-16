package webserver

import (
	"net/http"
)

func registerEndpoints(endpoints *Endpoints) {

	endpoints.add("/route-1", GET, func(request *http.Request) (StatusCode, string) {
		body := "Hello, You hit Route One!"
		return OK, body
	})

	endpoints.add("/route-1", POST, func(request *http.Request) (StatusCode, string) {
		buf := make([]byte, 1024)
		request.Body.Read(buf)
		body := "Hello, You hit Route One with content " + string(buf)
		return OK, body
	})

	endpoints.add("/route-2", GET, func(request *http.Request) (StatusCode, string) {
		body := "Hello, You hit Route Two!"
		return OK, body
	})
}
