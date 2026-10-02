package brain

import (
	"fmt"
	"strings"
)

func (b *Brain) StartSession() error {
	if err := b.stt.Start(); err != nil {
		return fmt.Errorf("failed to start speech recognition: %w", err)
	}

	defer b.stt.Close()

	for {
		fmt.Println("\nListening...")

		input, err := b.stt.Listen()
		if err != nil {
			return fmt.Errorf("speech recognition failed: %w", err)
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.EqualFold(input, "exit") {
			fmt.Println("Goodbye!")
			return nil
		}

		fmt.Println("You >", input)

		response, err := b.Run(input)
		if err != nil {
			return err
		}

		fmt.Println("Jarvis >", response)

		if err := b.tts.Speak(response); err != nil {
			return fmt.Errorf("text-to-speech failed: %w", err)
		}
	}
}