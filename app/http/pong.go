package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/response"
)

func Pong(c *gin.Context) {
	response.Success(c, "pong")
}
