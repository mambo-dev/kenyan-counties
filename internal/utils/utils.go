package utils

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
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

type LimitOffset struct {
	Limit  int32
	Offset int32
}

func SafeInt32(n int) (int32, error) {
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, fmt.Errorf("integer overflow: %d out of int32 range", n)
	}
	return int32(n), nil
}

func SafeLimitOffsetParser(r *http.Request, itemCount int64) (*LimitOffset, string, error) {
	var limit int64 = itemCount
	var offset int64 = 0

	if r.URL.Query().Get("limit") != "" {
		parsedLimit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)

		if err != nil {

			return nil, "Invalid limit param", err
		}

		limit = int64(parsedLimit)

	}

	if r.URL.Query().Get("offset") != "" {
		parsedOffset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)

		if err != nil {

			return nil, "Invalid offset param", err
		}

		offset = int64(parsedOffset)
	}

	safeLimit, err := SafeInt32(int(limit))

	if err != nil {

		return nil, "Invalid limit param", err
	}

	safeOffset, err := SafeInt32(int(offset))

	if err != nil {

		return nil, "Invalid offset param", err
	}

	return &LimitOffset{
		Limit:  safeLimit,
		Offset: safeOffset,
	}, "", nil
}
