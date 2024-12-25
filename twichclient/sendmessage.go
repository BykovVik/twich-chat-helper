package twichclient

import (
	"os"

	"github.com/gempir/go-twitch-irc/v4"
)

var messageSum int

func SendMessage(client *twitch.Client, maxMessageSum int) {

	messageSum++

	if messageSum == maxMessageSum {
		channel := os.Getenv("CHANNEL_NAME")
		message := os.Getenv("MESSAGE")
		client.Say(channel, message)
		messageSum = 0
	}
}
