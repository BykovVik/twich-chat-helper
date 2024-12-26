package twichclient

import (
	"log"
	"t_chat/sound"
)

func PlaySound() {
	err := sound.PlaySound("./sound/message.wav")
	if err != nil {
		log.Fatal("Audio playback error: ", err)
	}
}
