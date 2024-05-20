package main

import (
	"log"
	"net/http"
)

type inMemoryPlayerStore struct{}

func (i *inMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&inMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
