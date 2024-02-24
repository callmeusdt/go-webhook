package controller

import (
	"github.com/gin-gonic/gin"
	"go-webhook/http"
)

func Ping(c *gin.Context) {
	http.Success(c, "pong", "pong")
}
