package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hook/bootstrap"
	"log"
	"net/http"
	"os/exec"
)

type Payload struct {
	Ref string `json:"ref"`
}

func init() {
	bootstrap.LoadConfig()
	bootstrap.InitLogger()
}

func main() {
	// 定义一个 GET 请求的路由
	bootstrap.CTX.POST("/hook", func(c *gin.Context) {
		log.Printf("[%v][%v], req:%+v", c.Request.Method, c.Request.RemoteAddr, c.Request.Body)

		secret := c.Request.Header.Get("X-Hub-Signature")
		if secret == "" || bootstrap.CFG.App.Secret != secret {
			log.Printf("Auth failed, secret: %s", secret)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Auth",
			})
			return
		}

		var payload Payload
		if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
			log.Printf("%s: %v", http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		if payload.Ref != "refs/heads/master" {
			log.Println("Ignore all other branches")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Ignore all other branches",
			})
			return
		}

		var stderr bytes.Buffer
		for _, item := range bootstrap.CFG.Commands {
			cmd := exec.Command(item.Name, item.Args...)
			cmd.Stderr = &stderr

			if err := cmd.Run(); err != nil {
				log.Printf("exec %v failed: %v", item.Name, err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("%v", string(stderr.Bytes())),
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "done!",
		})
	})

	// 启动 HTTP 服务器
	log.Fatal(
		bootstrap.CTX.Run(fmt.Sprintf(":%s", bootstrap.CFG.App.Port)).Error(),
	)
}
