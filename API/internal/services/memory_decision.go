package services

import (
	"encoding/json"
	"fmt"
	"strings"

	"veerai/internal/ai"
	"veerai/internal/models"
)

type MemoryDecision struct {
	Remember bool   `json:"remember"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

func DecideMemory(provider ai.Provider, message string) (*MemoryDecision, error) {
	prompt := `Decide whether the following user message contains a useful long-term fact or preference that should be remembered.

Only remember information that would be useful in future conversations.

Examples worth remembering:
- name
- preferences
- important personal facts
- recurring preferences

Do not remember:
- temporary questions
- ordinary conversation
- calculations
- one-time requests

Return ONLY valid JSON in this exact format:

{
  "remember": true,
  "key": "name",
  "value": "Veer"
}

If nothing should be remembered, return:

{
  "remember": false,
  "key": "",
  "value": ""
}

USER MESSAGE:

` + message

	messages := []models.Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	response, err := provider.Chat(messages)
	if err != nil {
		return nil, err
	}

	var decision MemoryDecision

	response = strings.TrimSpace(response)
	response = strings.TrimPrefix(response, "```json")
	response = strings.TrimPrefix(response, "```")
	response = strings.TrimSuffix(response, "```")
	response = strings.TrimSpace(response)

	err = json.Unmarshal([]byte(response), &decision)
	if err != nil {
		return nil, fmt.Errorf("invalid memory decision: %w", err)
	}

	return &decision, nil
}
