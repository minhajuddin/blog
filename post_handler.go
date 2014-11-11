package main

import (
	"log"
	"net/http"
	"strconv"
)

func createPost(w http.ResponseWriter, r *http.Request) {
	//save
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err, r.Form)
		w.WriteHeader(500)
		return //set return to 500
	}

	p := Post{
		ID:    id,
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
	}
	p.save()
	//redirect
	http.Redirect(w, r, "/", http.StatusFound)
}
