package limiter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRateLimiting(t *testing.T) {
	// Use a more aggressive rate limiter for testing purposes
	rl := NewRateLimiter(rate.Every(10*time.Millisecond), 1)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	// Setup a test server with the Limit middleware
	testServer := httptest.NewServer(rl.Limit(handler))
	defer testServer.Close()

	client := &http.Client{}

	// Allow time for at least one request
	time.Sleep(10 * time.Millisecond)

	// This request should pass
	resp, err := client.Get(testServer.URL)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d for the allowed request", http.StatusOK, resp.StatusCode)
	}
	resp.Body.Close()

	// This request should be rate limited
	resp, err = client.Get(testServer.URL)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected status %d, got %d for the rate-limited request", http.StatusTooManyRequests, resp.StatusCode)
	}
}

func TestCleanupVisitors(t *testing.T) {
	rl := NewRateLimiter(rate.Limit(1), 3)

	// add a visitor last seen 4 mimutes ago
	rl.mu.Lock()
	rl.visitors["test-ip"] = &visitor{limiter: rate.NewLimiter(1, 3), lastSeen: time.Now().Add(-4 * time.Minute)}
	rl.mu.Unlock()

	// Trigger cleanup manually instead of waiting
	rl.CleanupVisitors()

	// Check if the visitor was cleaned up
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	if _, exists := rl.visitors["test-ip"]; exists {
		t.Error("expected visitor to be cleaned up, but it still exists")
	}
}
