package ai

import "veerai/internal/models"

type ToolCall struct {
	ID        string
	Name      string
	Arguments map[string]any
}

type ChatResult struct {
	Content   string
	ToolCalls []ToolCall
}

type Provider interface {
	Chat(messages []models.Message) (string, error)

	ChatWithTools(
		messages []models.Message,
		tools []ToolDefinition,
	) (*ChatResult, error)
}

type ToolDefinition struct {
	Name        string
	Description string
	Parameters  []Parameter
}

type Parameter struct {
	Name        string
	Type        string
	Description string
	Required    bool
}
