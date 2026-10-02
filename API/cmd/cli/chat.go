package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func sendChat(conversationID int64, token string, message string) {
	reqBody := ChatRequest{
		ConversationID: conversationID,
		Message:        message,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080/api/chat",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Could not connect to API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("API error:", string(body))
		return
	}

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("VeerAI >", result.Response)
}
