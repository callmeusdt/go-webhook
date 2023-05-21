package _default

import (
	"github.com/gin-gonic/gin"
	"hook/bootstrap"
)

func Pong(c *gin.Context) {
	bootstrap.ResponsePlain(c, "pong")
}
