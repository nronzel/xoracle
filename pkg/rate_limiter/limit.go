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

// Create a map to hold the rate limiters for each visitor and a mutex
var visitors = make(map[string]*visitor)
var mu sync.RWMutex

// Run a background goroutine to remove old entries from the visitors map to
// prevent unbounded growth.
func init() {
	go cleanupVisitors()
}

// Retrieve and return the rate limiter for the current user if it already
// exists. Otherwise create a new rate limiter and add it to the visitors map,
// using IP address as key.
func getVisitor(ip string) *rate.Limiter {
	mu.RLock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
        // NewLimiter(rate, burst)
		limiter := rate.NewLimiter(1, 3)
		// Include current time when creating a new visitor.
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}
	// Update the last seen time for visitor.
	v.lastSeen = time.Now()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for more
// than 3 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the IP for the current user
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Call getVisitor func to retrive the rate limiter for the current user
		limiter := getVisitor(ip)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
