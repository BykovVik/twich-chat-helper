package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"t_chat/sound"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()

	// base data block
	messageSum := 1
	maxMessageSum := os.Getenv("MAX_MESSAGE")
	maxNum, error_type := strconv.Atoi(maxMessageSum)

	if error_type != nil {
		fmt.Printf("Ошибка заполнения конфигурационного файла: %v\n", error_type)
	}

	//Twich client block
	client := twitch.NewClient("gypsy_studio", "oauth:"+fmt.Sprint(os.Getenv("ACCESS_TOKEN")))

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		playSound()

		if messageSum < maxNum {
			messageSum += 1
			return
		} else {
			sendMessage(client)
			messageSum = 0
		}

	})

	client.Join(os.Getenv("CHANNEL_NAME"))
	err := client.Connect()

	if err != nil {
		panic(err)
	}

}

func playSound() {
	err := sound.PlaySound("message.wav")
	if err != nil {
		fmt.Printf("Ошибка при воспроизведении звука: %v\n", err)
	}
}

func sendMessage(client *twitch.Client) {
	channel := os.Getenv("CHANNEL_NAME")
	message := os.Getenv("MESSAGE")
	client.Say(channel, message)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}
