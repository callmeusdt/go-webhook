package bootstrap

import (
	"fmt"
	"log"
	"os"
)

var LOGGER *os.File

func InitLogger() {
	LOGGER, err := os.OpenFile("./log/hook.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(LOGGER *os.File) {
		err := LOGGER.Close()
		if err != nil {
			fmt.Sprintf("logger closed unexpectedly")
		}
	}(LOGGER)
	log.SetOutput(LOGGER)
}
