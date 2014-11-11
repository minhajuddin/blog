package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPostSave(t *testing.T) {
	p := Post{ID: 1,
		Title: "Hello World!",
		Body:  "This is my first post!\nI am excited!!!\n\nThis is going to be awesome.",
	}
	p.save()

	//read it from the filesystem
	data, err := ioutil.ReadFile("./data/1")
	if err != nil {
		t.Errorf("file not saved %s", err)
	}

	s := string(data)
	if !strings.Contains(s, "Hello World!") {
		t.Errorf("title not present in: %s.", s)
	}

	if !strings.Contains(s, "This is my first post!") {
		t.Errorf("body not present in: %s.", s)
	}
}

func TestFindPost(t *testing.T) {
	p := Post{ID: 132,
		Title: "Archer",
		Body:  "Do you want ants?",
	}
	p.save()

	post := findPost(p.ID)
	if p.ID != post.ID || p.Title != post.Title || p.Body != post.Body {
		t.Errorf("expected %#v but got %#v", p, *post)
	}
}

func TestAllPosts(t *testing.T) {
	os.RemoveAll(dataDir)
	(&Post{ID: 12, Title: "AA", Body: "BB"}).save()
	(&Post{ID: 13, Title: "DD", Body: "EE"}).save()
	posts := allPosts()
	if len(posts) != 2 {
		t.Errorf("got %v but expected %v", len(posts), 2)
	}
	if posts[0].Title != "AA" || posts[0].Body != "BB" {
		t.Errorf("got '%v' but expected '%v'", posts[0].Title, "AA")
	}
	if posts[1].Title != "DD" || posts[1].Body != "EE" {
		t.Errorf("got %v.", posts[1].Title)
	}

}
