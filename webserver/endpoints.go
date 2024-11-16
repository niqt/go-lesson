package webserver

import (
	"net/http"
)

type HandlerFunc func(request *http.Request) (statusCode StatusCode, body string)

type Endpoints struct {
	endpointsMap map[string]map[HttpMethod]HandlerFunc
}

func NewEndpoints() *Endpoints {
	return &Endpoints{endpointsMap: map[string]map[HttpMethod]HandlerFunc{}}
}
func (e *Endpoints) add(route string, method HttpMethod, handler HandlerFunc) {
	methodToHandlerMap := e.endpointsMap[route]
	if len(methodToHandlerMap) == 0 {
		e.endpointsMap[route] = map[HttpMethod]HandlerFunc{}
	}
	e.endpointsMap[route][method] = handler
}
