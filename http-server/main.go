package main

import (
	"log"
	"net/http"
)

func main() {
	// playerServer := &PlayerServer{NewInMemoryPlayerStore()}
	postgresPlayerStore, err := NewPostgresPlayerStore()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to db")

	err = postgresPlayerStore.Init()
	if err != nil {
		log.Fatal(err)
	}

	playerServer := &PlayerServer{postgresPlayerStore}

	log.Fatal(http.ListenAndServe(":5000", playerServer))
}
