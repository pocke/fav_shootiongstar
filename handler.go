package main

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("hogefugastring"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("index")

	session, err := store.Get(r, "twitter")
	if err != nil {
		// TODO: error handling
	}

	id, ok := session.Values["id"].(string)
	// TODO
	if !(ok && userExists(id)) {
		log.Println("login")
	} else {
		log.Println("logined")
	}
}
