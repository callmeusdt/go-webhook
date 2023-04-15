package bootstrap

import (
	"log"
	"os"
)

func InitLogger() {
	logFile, err := os.OpenFile("./log/hook.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)
	log.SetOutput(logFile)
}
