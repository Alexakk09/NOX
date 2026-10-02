package services

import (
	"errors"
	"strings"

	"veerai/internal/ai"
	"veerai/internal/models"
	"veerai/internal/repository"
)

var ErrConversationForbidden = errors.New(
	"conversation does not belong to user",
)

type ChatService struct {
	provider ai.Provider
}

func NewChatService(provider ai.Provider) *ChatService {
	return &ChatService{
		provider: provider,
	}
}

func buildMemoryMessage(memories map[string]string) models.Message {
	memoryText := "Known facts about this user:\n"

	for key, value := range memories {
		memoryText += "- " + key + ": " + value + "\n"
	}

	return models.Message{
		Role:    "system",
		Content: memoryText,
	}
}

func addMemoryToMessages(
	messages []models.Message,
	memoryMessage models.Message,
) []models.Message {
	result := []models.Message{memoryMessage}
	result = append(result, messages...)
	return result
}

func convertToolDefinitions(
	definitions []models.ToolDefinition,
) []ai.ToolDefinition {
	result := make([]ai.ToolDefinition, 0, len(definitions))

	for _, definition := range definitions {
		parameters := make([]ai.Parameter, 0, len(definition.Parameters))

		for _, parameter := range definition.Parameters {
			parameters = append(parameters, ai.Parameter{
				Name:        parameter.Name,
				Type:        parameter.Type,
				Description: parameter.Description,
				Required:    parameter.Required,
			})
		}

		result = append(result, ai.ToolDefinition{
			Name:        definition.Name,
			Description: definition.Description,
			Parameters:  parameters,
		})
	}

	return result
}

func buildConversationTitle(message string) string {
	title := strings.Join(strings.Fields(message), " ")

	if title == "" {
		return "New Chat"
	}

	titleRunes := []rune(title)

	if len(titleRunes) > 60 {
		return string(titleRunes[:60]) + "..."
	}

	return title
}

func (s *ChatService) Chat(
	userID int64,
	conversationID int64,
	message string,
	toolDefinitions []models.ToolDefinition,
	toolResult *models.ToolResult,
) (*models.ChatResponse, error) {

	conversationUserID, err := repository.GetConversationUserID(conversationID)
	if err != nil {
		return nil, err
	}

	if conversationUserID != userID {
		return nil, ErrConversationForbidden
	}

	memories, err := GetUserMemories(userID)
	if err != nil {
		return nil, err
	}

	messages, err := repository.GetMessages(conversationID)
	if err != nil {
		return nil, err
	}

	isFirstMessage := len(messages) == 0

	if len(memories) > 0 {
		messages = addMemoryToMessages(
			messages,
			buildMemoryMessage(memories),
		)
	}

	

	if message != "" {
		userMessage := models.Message{
			Role:    "user",
			Content: message,
		}

		messages = append(messages, userMessage)

		if err := ProcessMemory(
			s.provider,
			userID,
			message,
		); err != nil {
			return nil, err
		}

		if err := repository.AddMessage(
			conversationID,
			userMessage,
		); err != nil {
			return nil, err
		}

		if isFirstMessage {
			title := buildConversationTitle(message)

			if err := repository.UpdateConversationTitle(
				conversationID,
				title,
			); err != nil {
				return nil, err
			}
		}
	}

	// When Jarvis sends a tool result, reconstruct the assistant
	// tool-call message followed by the tool result message.
	if toolResult != nil {
		messages = append(messages, models.Message{
			Role:      "assistant",
			Content:   "",
			ToolCalls: []models.ToolCall{toolResult.Call},
		})

		messages = append(messages, models.Message{
			Role:       "tool",
			Content:    toolResult.Result,
			ToolCallID: toolResult.Call.ID,
		})
	}

	result, err := s.provider.ChatWithTools(
		messages,
		convertToolDefinitions(toolDefinitions),
	)
	if err != nil {
		return nil, err
	}

	response := &models.ChatResponse{
		Response: result.Content,
	}

	for _, call := range result.ToolCalls {
		response.ToolCalls = append(
			response.ToolCalls,
			models.ToolCall{
				ID:        call.ID,
				Name:      call.Name,
				Arguments: call.Arguments,
			},
		)
	}

	// Only persist the final assistant response.
	// Tool-call intermediate state is reconstructed from the
	// ToolResult sent by Jarvis on the next request.
	if len(result.ToolCalls) == 0 {
		assistantMessage := models.Message{
			Role:    "assistant",
			Content: result.Content,
		}

		if err := repository.AddMessage(
			conversationID,
			assistantMessage,
		); err != nil {
			return nil, err
		}
	}

	return response, nil
}
