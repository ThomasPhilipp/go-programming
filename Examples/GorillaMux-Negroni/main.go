package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// An example of using a more feature rich router ("Gorilla/Mux") and
// middleware ("Negroni"). It also implements a custom middleware which can be
// called into a chain. By calling "http://localhost:8080/hello" you will get
// an unauthorized message - otherwise you have to call http://localhost:8080/hello?username=admin&password=password
//
// Source: github.com/blackhat-go/bhg/ch-4/negroni _example/

type trivial struct {
	// empty
}

func (t *trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Executing trivial middleware")
	next(w, r)
}

type BasicAuth struct {
	Username string
	Password string
}

func (b *BasicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Grabs username and password from the request
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// Compares username and password with the stored once
	if username != b.Username || password != b.Password {
		http.Error(w, "Unauthorized", 401)
		return
	}

	// Adding username to the request context for
	ctx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(ctx)

	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Reads the value from the request context to display it
	username := r.Context().Value("username").(string)
	// Writes the custom response
	fmt.Fprintf(w, "Hi %s\n", username)
}

func main() {
	// Creates a router
	r:= mux.NewRouter()

	r.HandleFunc("/hello", hello).Methods("GET")

	// Uses the classic middleware, including logging, recovery from panics
	// and serve files from the public folder
	n := negroni.Classic()

	// Uses customized middleware
	n.Use(&trivial{})
	n.Use(&BasicAuth{
		Username: "admin",
		Password: "password",
	})

	// Uses the router into the middleware stack
	n.UseHandler(r)

	http.ListenAndServe(":8080", n)
}