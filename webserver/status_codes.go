package webserver

type StatusCode string

const (
	OK               StatusCode = "200 OK"
	MethodNotAllowed StatusCode = "405 Method Not Allowed"
	NotFound         StatusCode = "404 Not Found"
)
