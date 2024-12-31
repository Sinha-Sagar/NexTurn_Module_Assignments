package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("%s - %s %s %d %s", c.ClientIP(), c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency)
	}
}