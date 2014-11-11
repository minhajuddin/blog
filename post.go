package main

type Post struct {
	ID    int
	Title string
	Body  string
}

func (p *Post) save() {
	//save this post
}
