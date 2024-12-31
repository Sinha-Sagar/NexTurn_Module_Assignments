package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	requests   = make(map[string]int)
	mu         sync.Mutex
	rateLimit  = 10
	rateWindow = time.Minute
)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		requests[ip]++
		count := requests[ip]
		if count == 1 {
			go func(ip string) {
				time.Sleep(rateWindow)
				mu.Lock()
				delete(requests, ip)
				mu.Unlock()
			}(ip)
		}
		mu.Unlock()

		if count > rateLimit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
