package main

type ChatRequest struct {
	ConversationID int64  `json:"conversation_id"`
	Message        string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

type PDFResponse struct {
	Message string `json:"message"`
	Text    string `json:"text"`
}

type PDFExplainResponse struct {
	Response string `json:"response"`
}

type PDFSummaryResponse struct {
	Response string `json:"response"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
