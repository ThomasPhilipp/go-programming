package main

import (
	"fmt"
	"net/http"
)

// Example of starting a server that handles requests to a single path with a
// customized response. The resource is available at "/hello?name=alice".
// The code is using the DefaultServerMux. A ServerMux is a server multiplexer,
// that can handle multiple HTTP requests by doing this using goroutines (1
// goroutine per incoming request)
//
// Source: https://github.com/blackhat-go/bhg/blob/master/ch-4/hello_world/main.go

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
	// Handle various resources
	http.HandleFunc("/hello", hello)

	// Starts ah HTTP server and listen on port 8080 across all interfaces
	http.ListenAndServe(":8080", nil) // nil = DefaultServerMux
}
