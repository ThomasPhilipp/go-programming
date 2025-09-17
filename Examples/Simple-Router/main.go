package main

import (
	"fmt"
	"net/http"
)

// Example of a router which dynamically handle inbound requests by inspecting
// the URL path. The implemented resources are: "/a", "/b", etc. If no matching
// resource is found, a "404" is returned.
//
// A more powerful package for routing is: "gorilla/mux", which provides also
// routing based on regex, parameter or verb matching as well as sub routing.
//
// Source: https://github.com/blackhat-go/bhg/tree/master/ch-4/simple_router

// The router type implements the http.Handler interface
type router struct {
	// empty
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprintf(w, "Executing /a")
	case "/b":
		fmt.Fprintf(w, "Executing /b")
	case "/c":
		fmt.Fprintf(w, "Executing /c")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router

	http.ListenAndServe(":8080", &r)
}