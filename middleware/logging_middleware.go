package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	// Using a anonymous function for the logging operation.
	return func(c *gin.Context) {
		
		startTime := time.Now()

		
		c.Next()

		
		endTime := time.Now()

		latency := endTime.Sub(startTime)

		
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		
		log.Printf("METHOD: %s | PATH: %s | STATUS: %d | LATENCY: %v", method, path, statusCode, latency)
	}
}
