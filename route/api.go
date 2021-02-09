package route

import (
	"github.com/gin-gonic/gin"

	"github.com/gopher-lego/ginger/app/http"
	"github.com/gopher-lego/ginger/app/middleware"
)

var corsMiddleware = middleware.CorsMiddleware()
var limiterMiddleware = middleware.LimiterMiddleware()

func Set(e *gin.Engine) {
	e.Use(corsMiddleware)
	e.Use(limiterMiddleware)

	// Top routes prefix
	top := e.Group("/api")

	heartbeat := top.Group("/ping")
	{
		heartbeat.GET("", http.Pong)
	}
}
