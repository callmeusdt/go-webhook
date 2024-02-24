package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取请求 Body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// 恢复请求 Body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("[REQ] %v %s @%v DATA:%s", c.Request.Method, c.Request.RequestURI, c.Request.RemoteAddr, string(body))

		c.Next()
	}
}
