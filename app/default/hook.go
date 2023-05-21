package _default

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hook/app"
	"hook/bootstrap"
	"hook/utils"
	"log"
	"net/http"
	"os/exec"
)

func Hook(c *gin.Context) {
	secret := c.Request.Header.Get("X-Hub-Signature")
	if secret == "" || !utils.VerifySecretToken(c.Request, app.CFG.App.Secret) {
		bootstrap.Fail(c, http.StatusBadRequest, "Auth failed")
		return
	}

	// 获取请求体中的内容
	_, err := c.GetRawData()
	if err != nil {
		bootstrap.Fail(c, http.StatusBadRequest, "Empty request")
		return
	}

	for _, item := range app.CFG.Commands {
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
}
