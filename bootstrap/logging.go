package bootstrap

import (
	"fmt"
	"log"
	"os"
)

func InitLogger(appName string) {
	LogFile, err := os.OpenFile(
		fmt.Sprintf("./log/%s.log", appName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(LogFile)
	//fmt.Println("logger inited")
}
