package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var SERVER *gin.Engine

func Success(c *gin.Context, data interface{}, msg string) {
	Json(c, 200, data, msg)
}

func Fail(c *gin.Context, code int, msg string) {
	Json(c, code, nil, msg)
}

func Json(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func InitLogger(name string) {
	dirname := "./log"
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		os.MkdirAll(dirname, 0755)
		fmt.Printf("Directory '%s' created\n", dirname)
	} else if err != nil {
		log.Fatalf("Error occurred while checking directory: %s, error: %w", dirname, err)
	}

	fileName := fmt.Sprintf("%s/%s.log", dirname, name)
	LogFile, err := os.OpenFile(
		fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(LogFile)
}
