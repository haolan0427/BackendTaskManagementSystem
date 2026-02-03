package middleware

import (
    "net/http"
    "sync"
    "time"
    
    "github.com/gin-gonic/gin"
)

type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.Mutex
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (r *RateLimiter) Allow(ip string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    now := time.Now()
    requests := r.requests[ip]
    
    var validRequests []time.Time
    for _, req := range requests {
        if now.Sub(req) < r.window {
            validRequests = append(validRequests, req)
        }
    }
    
    r.requests[ip] = validRequests
    
    if len(validRequests) >= r.limit {
        return false
    }
    
    r.requests[ip] = append(validRequests, now)
    return true
}

func RateLimitMiddleware(limiter *RateLimiter) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        if !limiter.Allow(ip) {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}