package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取请求 Body
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// 恢复请求 Body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		log.Printf("%v %s @%v DATA:%s", c.Request.Method, c.Request.RequestURI, c.Request.RemoteAddr, string(body))

		c.Next()
	}
}
