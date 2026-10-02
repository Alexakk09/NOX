package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ToolCall struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

type ChatResponse struct {
	Response  string     `json:"response"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

type ToolParameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type ToolDefinition struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Parameters  []ToolParameter `json:"parameters,omitempty"`
}

type ToolResult struct {
	Call   ToolCall `json:"call"`
	Result string   `json:"result"`
}

type ChatRequest struct {
	ConversationID  int64            `json:"conversation_id"`
	Message         string           `json:"message,omitempty"`
	ToolDefinitions []ToolDefinition `json:"tool_definitions,omitempty"`
	ToolResult      *ToolResult      `json:"tool_result,omitempty"`
}

type CreateConversationResponse struct {
	ID int64 `json:"id"`
}

func (c *APIClient) CreateConversation() (int64, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		c.BaseURL+"/api/conversations",
		nil,
	)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf(
			"create conversation failed: %s",
			string(body),
		)
	}

	var result CreateConversationResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.ID, nil
}

func (c *APIClient) SendChat(
	request ChatRequest,
) (*ChatResponse, error) {
	reqBody := request

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.BaseURL+"/api/chat",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(
		"Authorization",
		"Bearer "+c.Token,
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"chat failed: %s",
			string(body),
		)
	}

	var result ChatResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
