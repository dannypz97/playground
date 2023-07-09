// Token bucket algorithm implementation

package limiter

import (
	"sync"
	"time"
)

type Rate uint32

type RateLimiter struct {
	sync.Mutex
	rate       Rate
	burst      uint32 // max tokens
	tokenCount uint32
	last       time.Time //last time token field was updated
}

// Returns a new RateLimiter that allows events up to rate r and permits
// bursts of at most b tokens.
func NewRateLimiter(r Rate, b int) *RateLimiter {
	return &RateLimiter{
		rate:       r,
		burst:      uint32(b),
		tokenCount: uint32(b),
		last:       time.Now(),
	}
}

func (r *RateLimiter) IsAllowed() bool {
	r.Lock()
	defer r.Unlock()

	now := time.Now()

	// Calculate the number of tokens that should have been added since r.last
	delta := uint32(now.Sub(r.last).Seconds() * float64(r.rate))

	if r.tokenCount+delta < r.burst {
		r.tokenCount += delta
	} else {
		r.tokenCount = r.burst
	}
	r.last = now

	if r.tokenCount > 0 {
		r.tokenCount--
		return true
	}
	return false
}
