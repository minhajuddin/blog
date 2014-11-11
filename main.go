package main

import (
	"log"
	"net/http"
)

//port on which we start the server
const port = ":3000"

func main() {
	//a request for path beginning with slash (which is all paths)
	//should be handled by the hello function
	http.HandleFunc("/", hello)

	http.HandleFunc("/posts", createPost)

	log.Println("starting a server on http://localhost", port)
	//start the server on this port
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//the hello function is a typical web handler which has
//access to the incoming request r and can create a response
//using the ResponseWriter w
func hello(w http.ResponseWriter, r *http.Request) {
	//write "Hello World" to the response stream
	_, err := w.Write([]byte("Hello World!"))

	// if there is an error log it
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
