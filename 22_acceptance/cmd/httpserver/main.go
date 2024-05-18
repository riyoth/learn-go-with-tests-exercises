package main

import (
	"log"
	"net/http"

	"github.com/quii/go-specs-greet/adapters/httpserver"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/greet", http.HandlerFunc(httpserver.GreetHandler))
	mux.Handle("/curse", http.HandlerFunc(httpserver.CurseHandler))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
