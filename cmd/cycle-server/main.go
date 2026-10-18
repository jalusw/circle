package main

import (
	"log"
	"net/http"

	"github.com/jalusw/cycle/internal/api"
	"github.com/jalusw/cycle/internal/store"
)

func main() {
	s, err := store.Open("cycle.db")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	if err := s.Migrate(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	api.New(mux).Routes()
	log.Println("cycle-server: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
