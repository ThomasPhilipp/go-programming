package main

import (
	"fmt"
	"log"
	"net/http"
)

// Example of a middleware, which execute on all incoming requests regardless
// of the destination function. In that case the start and stop time is logged.
// Other examples of using a middleware is for authenticating and authorizing
// users and mapping resources. An alternative package is "negroni".

// The logger type implements the http.Handler interface but also contains an
// Inner, which is a http.Handler itself
type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested:", r.URL.Path)
	l.Inner.ServeHTTP(w, r)
	log.Println("Finished")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {
	f := http.HandlerFunc(hello)

	// Creates the logger middleware
	l := logger{Inner: f}

	// Starts the server with a reference to the logger
	http.ListenAndServe(":8080", &l)
}