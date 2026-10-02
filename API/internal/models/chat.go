package models

type ChatRequest struct {
	ConversationID  int64            `json:"conversation_id"`
	Message         string           `json:"message,omitempty"`
	ToolDefinitions []ToolDefinition `json:"tool_definitions,omitempty"`
	ToolResult      *ToolResult      `json:"tool_result,omitempty"`
}

type ChatResponse struct {
	Response  string     `json:"response"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

type ToolDefinition struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type ToolCall struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

type ToolResult struct {
	Call ToolCall `json:"call"`
	Result string  `json:"result"`
}

type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}