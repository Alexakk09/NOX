package ai

import "veerai/internal/models"

type MockProvider struct{}

func (MockProvider) Chat(messages []models.Message) (string, error) {
	if len(messages) == 0 {
		return "Mock AI: no messages received", nil
	}

	lastMessage := messages[len(messages)-1]

	return "Mock AI response to: " + lastMessage.Content, nil
}