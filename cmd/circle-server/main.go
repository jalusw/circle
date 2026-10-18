package main

import (
	"log"
	"net/http"

	"github.com/jalusw/circle/internal/api"
	"github.com/jalusw/circle/internal/store"
)

func main() {
	s, err := store.Open("circle.db")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	if err := s.Migrate(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	api.New(mux).Routes()
	log.Println("circle-server: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
