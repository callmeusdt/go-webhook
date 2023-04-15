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
	bootstrap.GIN.POST("/hook", func(c *gin.Context) {
		log.Printf("[%v][%v], req:%+v", c.Request.Method, c.Request.RemoteAddr, c.Request.Body)

		secret := c.Request.Header.Get("X-Hub-Signature")
		if secret == "" || bootstrap.CFG.App.Secret != secret {
			bootstrap.Fail(c, http.StatusBadRequest, "Auth failed")
			return
		}

		// 获取请求体中的内容
		_, err := c.GetRawData()
		if err != nil {
			bootstrap.Fail(c, http.StatusBadRequest, "Empty request")
			return
		}

		for _, item := range bootstrap.CFG.Commands {
			cmd := exec.Command(item.Name, item.Args...)
			output, err := cmd.CombinedOutput()
			log.Printf("[%v] output: %v", item.Name, string(output))

			if err != nil {
				log.Printf("[%v] failed: %v", item.Name, err.Error())
				bootstrap.Fail(c, http.StatusInternalServerError, fmt.Sprintf("%v", string(output)))
				return
			}
		}

		bootstrap.Success(c, nil, "done!")
	})

	// 启动 HTTP 服务器
	log.Fatal(
		bootstrap.GIN.Run(fmt.Sprintf(":%s", bootstrap.CFG.App.Port)).Error(),
	)
}
