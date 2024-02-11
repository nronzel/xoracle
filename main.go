package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nronzel/xoracle/pkg/handlers"
	limiter "github.com/nronzel/xoracle/pkg/rate_limiter"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HandlerRoot)
	mux.HandleFunc("POST /decrypt", handlers.HandlerDecrypt)

	rl := limiter.NewRateLimiter(1, 3)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      rl.Limit(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("Server starting on port: 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Problem starting server: %v", err)
	}
}
