package bootstrap

import (
	"fmt"
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
		"message": obj.Message,
		"code":    obj.Code,
		"data":    obj.Data,
	})

	err := LOGGER.Sync()
	if err != nil {
		fmt.Sprintf("logger.sync failed: %v", err.Error())
		return
	}
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
