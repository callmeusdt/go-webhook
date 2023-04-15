package bootstrap

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var GIN = gin.Default()

type Response struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Data    *[]interface{} `json:"data"`
}

func ResponseJson(c *gin.Context, status int, obj Response) {
	log.Printf("response: %+v", obj)
	c.JSON(status, gin.H{
		"Message": obj.Message,
		"Code":    obj.Code,
		"Data":    obj.Data,
	})

	LOGGER.Sync()
}

func Success(c *gin.Context, data *[]interface{}, message string) {
	ResponseJson(c, http.StatusOK, Response{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	})
}

func Fail(c *gin.Context, status int, message string) {
	log.Printf("response:failed %s", message)
	ResponseJson(c, status, Response{
		Message: message,
		Code:    status,
	})
}
