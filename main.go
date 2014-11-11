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
	http.HandleFunc("/", home)

	http.HandleFunc("/posts", createPost)

	log.Println("starting a server on http://localhost", port)
	//start the server on this port
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
