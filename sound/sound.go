package sound

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PlaySound(soundFile string) error {

	if _, err := os.Stat(soundFile); os.IsNotExist(err) {
		return fmt.Errorf("файл %s не найден", soundFile)
	}

	var cmd *exec.Cmd

	// Определяем команду в зависимости от операционной системы
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("aplay", soundFile)
	case "darwin":
		cmd = exec.Command("afplay", soundFile)
	case "windows":
		cmd = exec.Command("powershell", "-c", "Start-Process", "powershell", "-ArgumentList", "-c", fmt.Sprintf("Add-Type –AssemblyName presentationCore;[System.Media.SoundPlayer]::new('%s').PlaySync();", soundFile))
	default:
		return fmt.Errorf("операционная система не поддерживается")
	}

	return cmd.Run()
}
