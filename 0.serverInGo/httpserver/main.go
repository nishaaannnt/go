package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	fmt.Println("got request", hasFirst)
	io.WriteString(w, "this is my website")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got hello request")
	io.WriteString(w, "Hello http\n")
}

func main() {
	// api routes
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	// starting the server on port 3333
	fmt.Println("running")
	// nil is saying we are using default mux
	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("runnning")
	}
}
