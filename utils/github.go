package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
)

func VerifySecretToken(req *http.Request, secretToken string) bool {
	// 从Header中获取X-Hub-Signature
	signature := req.Header.Get("X-Hub-Signature")

	// 读取请求体内容
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return false
	}

	// 使用secretToken对请求体进行HMAC SHA1加密
	mac := hmac.New(sha1.New, []byte(secretToken))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	// 比较生成的加密签名和Header中的X-Hub-Signature
	return signature == "sha1="+expectedMAC
}
