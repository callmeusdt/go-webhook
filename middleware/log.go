package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("REQUEST:METHOD:%s: URI:%s data:%+v", c.Request.Method, c.Request.RequestURI, c.Request.Body)
		c.Next()
	}
}
