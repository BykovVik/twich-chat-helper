package main

import (
	"fmt"
	"t_chat/twichclient"

	"github.com/joho/godotenv"
)

func main() {

	//doenv conf
	loadEnv()

	//start chat listener
	twichclient.StartTwichClient()

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Ошибка заполнения поля MAX_MESSAGE: %v\n", err)
	}
}
