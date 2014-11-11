package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	os.RemoveAll(dataDir)

	(&Post{ID: 12, Title: "AA", Body: "BB"}).save()
	(&Post{ID: 13, Title: "DD", Body: "EE"}).save()

	w := httptest.NewRecorder()
	home(w, nil)

	if w.Code != 200 {
		t.Errorf("got %v but expected %v", w.Code, 200)
	}

	if w.Header().Get("Content-Type") != "text/plain" {
		t.Errorf("got %v but expected %v", w.Header().Get("Content-Type"), "text/plain")
	}

	body := w.Body.String()

	if body != "12. AA\nBB\n\n13. DD\nEE\n\n" {
		t.Errorf("got %v.", body)
	}
}
