package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 自定义响应写入器
type responseWriter struct {
	gin.ResponseWriter
	Body []byte
}

// 重写 Write 方法，将响应体写入自定义缓冲区
func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.Body = append(rw.Body, data...)
	return rw.ResponseWriter.Write(data)
}

func LogResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := &responseWriter{ResponseWriter: c.Writer, Body: []byte{}}
		c.Writer = writer

		c.Next()
		//fmt.Println(writer.Status())
		log.Printf("[RESP] HTTP:%d BODY:%s", writer.Status(), string(writer.Body))
	}
}
