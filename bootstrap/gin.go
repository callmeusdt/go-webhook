package bootstrap

import (
	"github.com/gin-gonic/gin"
	"hook/middleware"
	"net/http"
)

type Response struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Data    *[]interface{} `json:"data"`
}

func InitGin() *gin.Engine {
	//gin.DefaultWriter = LogFile

	c := gin.Default()
	c.Use(middleware.LogRequest())
	c.Use(middleware.LogResponse())

	return c
}

func ResponseJson(c *gin.Context, status int, obj Response) {
	c.JSON(status, gin.H{
		"message": obj.Message,
		"code":    obj.Code,
		"data":    obj.Data,
	})
}

func ResponsePlain(c *gin.Context, content string) {
	c.String(200, content)
}

func Success(c *gin.Context, data *[]interface{}, message string) {
	ResponseJson(c, http.StatusOK, Response{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	})
}

func Fail(c *gin.Context, status int, message string) {
	ResponseJson(c, status, Response{
		Message: message,
		Code:    status,
	})
}
