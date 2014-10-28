package main

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/get_token", twitterGetTokenHandler)
	http.HandleFunc("/callback", twitterCallbackHandler)

	log.Println("start web server")
	log.Fatal(http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)))
}
