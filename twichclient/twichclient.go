package twichclient

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gempir/go-twitch-irc/v4"
)

func StartTwichClient() {

	//check type MAX_MESSAGE
	maxMessageSum, maxNumError := strconv.Atoi(os.Getenv("MAX_MESSAGE"))
	if maxNumError != nil {
		fmt.Printf("Ошибка заполнения поля MAX_MESSAGE: %v\n", maxNumError)
	}

	//create twich client
	client := twitch.NewClient(fmt.Sprint(os.Getenv("USERNAME")), "oauth:"+fmt.Sprint(os.Getenv("ACCESS_TOKEN")))

	//message handler
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		PlaySound()
		SendMessage(client, maxMessageSum)
	})

	client.Join(os.Getenv("CHANNEL_NAME"))
	err := client.Connect()

	if err != nil {
		panic(err)
	}
}
