package main

import (
	"log"
	"os"
	"t_chat/twichclient"

	"github.com/joho/godotenv"
)

func main() {
	//log settings
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Log file could not be opened: ", err)
	}

	defer logFile.Close()
	log.SetOutput(logFile)

	//load base data
	loadEnv()

	//start chat helper pacage
	twichclient.StartTwichClient()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file upload error: ", err)
	}
}
