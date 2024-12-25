package twichclient

import (
	"fmt"
	"t_chat/sound"
)

func PlaySound() {
	err := sound.PlaySound("message.wav")
	if err != nil {
		fmt.Printf("Ошибка при воспроизведении звука: %v\n", err)
	}
}
