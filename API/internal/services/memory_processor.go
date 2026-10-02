package services

import (
	"veerai/internal/ai"
)

func ProcessMemory(
	provider ai.Provider,
	userID int64,
	message string,
) error {
	decision, err := DecideMemory(provider, message)
	if err != nil {
		return err
	}

	if !decision.Remember {
		return nil
	}

	return SaveUserMemory(
		userID,
		decision.Key,
		decision.Value,
	)
}