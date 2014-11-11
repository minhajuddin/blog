package main

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("starting a server on http://localhost", port)
	http.ListenAndServe(port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
