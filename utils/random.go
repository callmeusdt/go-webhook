package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func RandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Sprintf("error: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)[:length]
}
