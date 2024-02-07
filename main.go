package main

import (
	"log"
)

func main() {

	// setup postgresStore
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	// start the server
	server := NewAPIServer(":3000", store)
	server.Run()
}
