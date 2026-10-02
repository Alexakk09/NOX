package services

import (
	"veerai/internal/repository"
)

func SaveUserMemory(userID int64, key string, value string) error {
	if userID <= 0 {
		return nil
	}

	if key == "" || value == "" {
		return nil
	}

	return repository.SaveMemory(userID, key, value)
}

func GetUserMemories(userID int64) (map[string]string, error) {
	if userID <= 0 {
		return nil, nil
	}

	return repository.GetMemories(userID)
}