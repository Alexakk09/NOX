package voice

import (
	"fmt"
	"os/exec"
)

type TTS struct{}

func NewTTS() *TTS {
	return &TTS{}
}

func (t *TTS) Speak(text string) error {
	if text == "" {
		return nil
	}

	script := fmt.Sprintf(`
Add-Type -AssemblyName System.Speech
$speaker = New-Object System.Speech.Synthesis.SpeechSynthesizer
$speaker.Speak(%q)
`, text)

	cmd := exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		script,
	)

	return cmd.Run()
}