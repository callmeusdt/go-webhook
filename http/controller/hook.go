package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-webhook/common"
	myhttp "go-webhook/http"
	"go-webhook/utils"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func Hook(c *gin.Context) {
	secret := c.Request.Header.Get("X-Hub-Signature")
	if secret == "" || !utils.VerifySecretToken(c.Request, common.GithubSecret) {
		myhttp.Fail(c, http.StatusBadRequest, "Auth failed")
		return
	}

	// 获取请求体中的内容
	_, err := c.GetRawData()
	if err != nil {
		myhttp.Fail(c, http.StatusBadRequest, "Empty request")
		return
	}

	cmd := exec.Command(common.HookCommand)
	pipeOut, _ := cmd.StdoutPipe()

	err = cmd.Start()
	if err != nil {
		log.Printf("[%v] failed: %v", common.HookCommand, err.Error())
		myhttp.Fail(c, http.StatusInternalServerError, fmt.Sprintf("%v", string(err.Error())))

		return
	}

	go func(cmd *exec.Cmd, pipeOut io.ReadCloser) {
		err = cmd.Wait()
		if err != nil {
			log.Printf("[%v] failed: %v", common.HookCommand, err.Error())
			return
		}

		output, _ := io.ReadAll(pipeOut)
		log.Printf("[%v] output: %+v", common.HookCommand, string(output))
	}(cmd, pipeOut)

	myhttp.Success(c, nil, "done!")
}
