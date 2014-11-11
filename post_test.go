package main

import (
	"io/ioutil"
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
		t.Errorf("title not present in:", s)
	}

	if !strings.Contains(s, "This is my first post!") {
		t.Errorf("body not present in:", s)
	}
}
