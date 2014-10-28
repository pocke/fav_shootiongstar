package main

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("start web server")
	log.Fatal(http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)))
}
