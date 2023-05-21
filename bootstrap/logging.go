package bootstrap

import (
	"fmt"
	"log"
	"os"
)

func InitLogger(appName string) {
	dirname := "./log"
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		os.MkdirAll(dirname, 0755)
		fmt.Printf("Directory '%s' created\n", dirname)
	} else if err != nil {
		log.Fatalf("Error occurred while checking directory: %s, error: %w", dirname, err)
	}

	fileName := fmt.Sprintf("%s/%s.log", dirname, appName)
	LogFile, err := os.OpenFile(
		fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(LogFile)
	//fmt.Println("logger inited")
}
