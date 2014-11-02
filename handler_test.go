package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: session test

func TestRootHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer ts.Close()

	res, _ := http.Get(ts.URL)

	if res.StatusCode != http.StatusOK {
		t.Error("status code error")
		return
	}
}

func TestSignoutHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(signoutHandler))
	defer ts.Close()

	res, _ := http.Get(ts.URL)
	if res.StatusCode != http.StatusSeeOther {
		t.Error("status code error")
		return
	}
}
