package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestPostCreateHandler(t *testing.T) {
	err := os.RemoveAll("./data")
	if err != nil {
		log.Println("unable to clean data", err)
	}
	w := httptest.NewRecorder()
	form := url.Values{"id": {"333"}, "title": {"God is an astronaut"}, "body": {"All is Violent, All is Bright"}}
	u, _ := url.Parse("http://example.com/posts")
	r := http.Request{
		URL:    u,
		Method: "POST",
		Body:   ioutil.NopCloser(strings.NewReader(form.Encode())),
		Header: http.Header{"Content-Type": {"application/x-www-form-urlencoded"}},
	}
	createPost(w, &r)

	//should return a 301 redirect
	if 302 != w.Code {
		t.Errorf("expected %v but got %v", 302, w.Code)
	}

	//should return a 301 redirect
	if "/" != w.Header().Get("Location") {
		t.Errorf("expected %v but got %v", "/", w.Header().Get("Location"))
	}

	//should save the file
	p := findPost(333)
	if p.Title != "God is an astronaut" {
		t.Errorf("got %v", p.Title)
	}

	if p.Body != "All is Violent, All is Bright" {
		t.Errorf("got %v", p.Body)
	}
}
