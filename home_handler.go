package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	posts := allPosts()
	for _, post := range posts {
		fmt.Fprintf(w, "%v. %s\n%s\n\n", post.ID, post.Title, post.Body)
	}
}
