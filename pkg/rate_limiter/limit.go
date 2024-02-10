package limiter

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// |~------------------~|
// |Per IP rate limiting|
// |~------------------~|

// Create custom visitor struct which holds the rate limiter for each visitor
// and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     rate.Limit
	burst    int
}

func NewRateLimiter(rate rate.Limit, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
	}
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			rl.CleanupVisitors()
		}
	}()
	return rl
}

// Retrieve and return the rate limiter for the current user if it already
// exists. Otherwise create a new rate limiter and add it to the visitors map,
// using IP address as key.
func (rl *RateLimiter) getVisitor(ip string) *rate.Limiter {
	rl.mu.RLock()
	v, exists := rl.visitors[ip]
	rl.mu.RUnlock()

	if exists {
		return v.limiter
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()
	// check if user was added while aquiring lock
	v, exists = rl.visitors[ip]
	if !exists {
		// NewLimiter(rate, burst)
		limiter := rate.NewLimiter(rl.rate, rl.burst)
		// Include current time when creating a new visitor.
		rl.visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}
	// Update the last seen time for visitor.
	v.lastSeen = time.Now()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for more
// than 3 minutes and delete the entries.
func (rl *RateLimiter) CleanupVisitors() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for ip, v := range rl.visitors {
		if time.Since(v.lastSeen) > 3*time.Minute {
			delete(rl.visitors, ip)
		}
	}
}

// Wraps HTTP handler with rate limiting.
func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the IP for the current user
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Call getVisitor func to retrive the rate limiter for the current user
		limiter := rl.getVisitor(ip)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
