package main

import (
	"fmt"
	"net/http"
)

//port on which we start the server
const port = ":3000"

func main() {
	//a request for path beginning with slash (which is all paths)
	//should be handled by the hello function
	http.HandleFunc("/", hello)
	fmt.Println("starting a server on http://localhost", port)
	//start the server on this port
	http.ListenAndServe(port, nil)
}

//the hello function is a typical web handler which has
//access to the incoming request r and can create a response
//using the ResponseWriter w
func hello(w http.ResponseWriter, r *http.Request) {
	//write "Hello World" to the response stream
	w.Write([]byte("Hello World!"))
}
