package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// "github.com/gin-contrib/cors"

	limiter "github.com/ulule/limiter/v3"
	limiterGinMw "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}

func LimiterMiddleware() gin.HandlerFunc {
	// Limiter middleware options
	limiterRate := limiter.Rate{
		Period: 60 * time.Second,
		Limit:  60,
	}
	// Current used only support single machine
	limiterStore := memory.NewStore()
	limiterInstance := limiter.New(limiterStore, limiterRate)

	return limiterGinMw.NewMiddleware(limiterInstance)
}
