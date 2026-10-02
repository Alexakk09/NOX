package services

import (
	"fmt"
	"os/exec"
	"veerai/internal/ai"
	"veerai/internal/models"
)

func ExtractPDFText(filePath string) (string, error) {
	cmd := exec.Command("pdftotext", "-layout", filePath, "-")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to extract PDF text: %w", err)
	}

	return string(output), nil
}

func SummarizePDF(provider ai.Provider, text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("PDF contains no extractable text")
	}

	prompt := `Summarize the following PDF clearly and concisely.

Include:
- the main topic
- the most important ideas
- important facts or conclusions
- key sections or concepts

Do not add information that is not present in the PDF.

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
