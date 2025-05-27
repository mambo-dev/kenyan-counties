package utils

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	requestsPerSecond = 5
	burstSize         = 3
	cleanupInterval   = time.Minute * 10
	staleAfter        = 10 * time.Minute
)

type ipLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	ipLimiters = make(map[string]*ipLimiter)
	mu         sync.Mutex
)

func GetLimiterForIP(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	lim, exists := ipLimiters[ip]
	if !exists {
		lim = &ipLimiter{
			limiter:  rate.NewLimiter(rate.Limit(requestsPerSecond), burstSize),
			lastSeen: time.Now(),
		}
		ipLimiters[ip] = lim
	} else {
		lim.lastSeen = time.Now()
	}

	return lim.limiter
}

func CleanUpStaleTimers() {
	ticker := time.NewTicker(cleanupInterval)
	for range ticker.C {
		mu.Lock()
		for ip, lim := range ipLimiters {
			if time.Since(lim.lastSeen) > staleAfter {
				delete(ipLimiters, ip)
			}
		}
		mu.Unlock()
	}
}
