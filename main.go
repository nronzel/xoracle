package main

import (
	"github.com/nronzel/xoracle/pkg/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.HandlerRoot)

	r.Post("/decrypt", handlers.HandlerDecrypt)

	log.Println("Server starting on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
