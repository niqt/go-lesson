package webserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

var endpoints *Endpoints

func ServerHttp() {
	// Create a TCP listener on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Printf("Failed to close listener: %v", err)
		}
	}(listener)
	log.Println("Server listening on :8080")

	// Initialize empty endpoints and register endpoints
	endpoints = NewEndpoints()
	registerEndpoints(endpoints)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Failed to close connection: %v", err)
		}
	}(conn)

	// Read the HTTP request from the connection
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Printf("Failed to read request: %v", err)
		return
	}

	// Print the request method and path
	log.Printf("Request: %s %s\n", request.Method, request.URL.Path)
	requestHandler(conn, request)
}

func requestHandler(conn net.Conn, request *http.Request) {
	methodToHandlerMap := endpoints.endpointsMap[request.URL.Path]
	if methodToHandlerMap == nil {
		writeResponse(conn, NotFound, "")
	}

	handler := methodToHandlerMap[HttpMethod(request.Method)]
	if handler == nil {
		writeResponse(conn, MethodNotAllowed, "")
	}

	statusCode, body := handler(request)
	writeResponse(conn, statusCode, body)
}

func writeResponse(conn net.Conn, statusCode StatusCode, response string) {
	response += "\r\n"
	_, err := conn.Write([]byte("HTTP/1.1 " + statusCode + "\r\n"))
	if err != nil {
		return
	}
	_, err = conn.Write([]byte("Content-Length: " + fmt.Sprint(len(response)) + "\r\n"))
	if err != nil {
		return
	}
	_, err = conn.Write([]byte("Content-Type: text/plain\r\n\r\n"))
	if err != nil {
		return
	}
	_, err = conn.Write([]byte(response))
	if err != nil {
		return
	}
}
