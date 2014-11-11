package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//unit test
func TestHello(t *testing.T) {
	r := httptest.NewRecorder()
	hello(r, nil)

	if r.Code != http.StatusOK {
		t.Errorf("incorrect response code %s", r.Code)
	}

	b := r.Body.String()
	if b != "Hello World!" {
		t.Errorf("incorrect response %s", b)
	}
}
