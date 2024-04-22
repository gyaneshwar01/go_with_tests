package main

import (
	"log"
	"net/http"
)

type InMemoryStore struct{}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryStore) RecordWin(name string) {}

func main() {
	store := InMemoryStore{}
	playerServer := &PlayerServer{&store}
	log.Fatal(http.ListenAndServe(":5000", playerServer))
}
