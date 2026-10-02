package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"veerai/internal/models"
)

type LMStudioProvider struct {
	BaseURL string
	Model   string
}

type lmStudioRequest struct {
	Model       string              `json:"model"`
	Messages    []lmStudioMessage   `json:"messages"`
	Temperature float64             `json:"temperature"`
	Tools       []lmStudioTool      `json:"tools,omitempty"`
}

type lmStudioMessage struct {
	Role       string             `json:"role"`
	Content    string             `json:"content"`
	ToolCalls  []lmStudioToolCall `json:"tool_calls,omitempty"`
	ToolCallID string             `json:"tool_call_id,omitempty"`
}

type lmStudioResponse struct {
	Choices []struct {
		Message struct {
			Role      string             `json:"role"`
			Content   string             `json:"content"`
			ToolCalls []lmStudioToolCall `json:"tool_calls"`
		} `json:"message"`
	} `json:"choices"`
}

type lmStudioTool struct {
	Type     string            `json:"type"`
	Function lmStudioFunction `json:"function"`
}

type lmStudioFunction struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Parameters  lmStudioParameters  `json:"parameters"`
}

type lmStudioParameters struct {
	Type       string                         `json:"type"`
	Properties map[string]lmStudioProperty    `json:"properties"`
	Required   []string                       `json:"required,omitempty"`
}

type lmStudioProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type lmStudioToolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type,omitempty"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

func NewLMStudioProvider(baseURL, model string) *LMStudioProvider {
	return &LMStudioProvider{
		BaseURL: baseURL,
		Model:   model,
	}
}

func convertMessages(messages []models.Message) []lmStudioMessage {
	result := make([]lmStudioMessage, 0, len(messages))

	for _, message := range messages {
		out := lmStudioMessage{
			Role:       message.Role,
			Content:    message.Content,
			ToolCallID: message.ToolCallID,
		}

		for _, call := range message.ToolCalls {
			arguments, err := json.Marshal(call.Arguments)
			if err != nil {
				continue
			}

			out.ToolCalls = append(out.ToolCalls, lmStudioToolCall{
				ID:   call.ID,
				Type: "function",
				Function: struct {
					Name      string `json:"name"`
					Arguments string `json:"arguments"`
				}{
					Name:      call.Name,
					Arguments: string(arguments),
				},
			})
		}

		result = append(result, out)
	}

	return result
}

func (p *LMStudioProvider) Chat(
	messages []models.Message,
) (string, error) {
	result, err := p.ChatWithTools(messages, nil)
	if err != nil {
		return "", err
	}

	return result.Content, nil
}

func convertToolDefinition(def ToolDefinition) lmStudioTool {
	properties := make(map[string]lmStudioProperty)
	required := []string{}

	for _, param := range def.Parameters {
		properties[param.Name] = lmStudioProperty{
			Type:        param.Type,
			Description: param.Description,
		}

		if param.Required {
			required = append(required, param.Name)
		}
	}

	return lmStudioTool{
		Type: "function",
		Function: lmStudioFunction{
			Name:        def.Name,
			Description: def.Description,
			Parameters: lmStudioParameters{
				Type:       "object",
				Properties: properties,
				Required:   required,
			},
		},
	}
}

func (p *LMStudioProvider) ChatWithTools(
	messages []models.Message,
	tools []ToolDefinition,
) (*ChatResult, error) {

	llmTools := make([]lmStudioTool, 0, len(tools))

	for _, tool := range tools {
		llmTools = append(
			llmTools,
			convertToolDefinition(tool),
		)
	}

	requestBody := lmStudioRequest{
		Model:       p.Model,
		Messages:    convertMessages(messages),
		Temperature: 0.7,
		Tools:       llmTools,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		p.BaseURL+"/chat/completions",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"LM Studio returned status %s",
			resp.Status,
		)
	}

	var result lmStudioResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Choices) == 0 {
		return nil, fmt.Errorf(
			"LM Studio returned no response",
		)
	}

	message := result.Choices[0].Message

	chatResult := &ChatResult{
		Content: message.Content,
	}

	for _, call := range message.ToolCalls {
		var arguments map[string]any

		if err := json.Unmarshal(
			[]byte(call.Function.Arguments),
			&arguments,
		); err != nil {
			return nil, fmt.Errorf(
				"invalid tool arguments for %s: %w",
				call.Function.Name,
				err,
			)
		}

		chatResult.ToolCalls = append(
			chatResult.ToolCalls,
			ToolCall{
				ID:        call.ID,
				Name:      call.Function.Name,
				Arguments: arguments,
			},
		)
	}

	return chatResult, nil
}