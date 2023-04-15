package main

import (
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

		// 获取请求体中的内容
		_, err := c.GetRawData()
		if err != nil {
			log.Printf("%s: %v", http.StatusBadRequest, err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}

		for _, item := range bootstrap.CFG.Commands {
			cmd := exec.Command(item.Name, item.Args...)
			output, err := cmd.CombinedOutput()
			log.Printf("[%v] output: %v", item.Name, string(output))

			if err != nil {
				log.Printf("[%v] failed: %v", item.Name, err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("%v", string(output)),
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
