package sound

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func PlaySound(soundFile string) error {

	//check sound file
	if _, err := os.Stat(soundFile); os.IsNotExist(err) {
		log.Fatal("Audio file not found: ", err)
	}

	var cmd *exec.Cmd

	//check OS
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("aplay", soundFile)
	case "darwin":
		cmd = exec.Command("afplay", soundFile)
	case "windows":
		cmd = exec.Command("powershell", "-c", "Start-Process", "powershell", "-ArgumentList", "-c", fmt.Sprintf("Add-Type –AssemblyName presentationCore;[System.Media.SoundPlayer]::new('%s').PlaySync();", soundFile))
	default:
		log.Fatal("Operating System Not Supported")
	}

	//play a beep
	return cmd.Run()
}
