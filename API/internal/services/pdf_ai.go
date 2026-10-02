package services

import (
	"fmt"

	"veerai/internal/ai"
	"veerai/internal/models"
)

func ExplainPDF(provider ai.Provider, text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("PDF contains no extractable text")
	}

	prompt := `Explain the following PDF content clearly and in detail.
Focus on the important concepts, definitions, and ideas.
Do not assume the reader already understands the material.

PDF CONTENT:

` + text

	messages := []models.Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	return provider.Chat(messages)
}
