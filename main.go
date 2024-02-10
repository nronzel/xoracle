package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nronzel/xoracle/pkg/handlers"
	limiter "github.com/nronzel/xoracle/pkg/rate_limiter"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.HandlerRoot)

	r.Post("/decrypt", handlers.HandlerDecrypt)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      limiter.Limit(r),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("Server starting on port: 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Problem starting server: %v", err)
	}
}
