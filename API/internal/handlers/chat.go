package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"veerai/internal/middleware"
	"veerai/internal/models"
	"veerai/internal/services"
)

type ChatHandler struct {
	service *services.ChatService
}

func NewChatHandler(service *services.ChatService) *ChatHandler {
	return &ChatHandler{
		service: service,
	}
}

func (h *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Only POST requests are allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var req models.ChatRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID <= 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response, err := h.service.Chat(
		userID,
		req.ConversationID,
		req.Message,
		req.ToolDefinitions,
		req.ToolResult,
	)

	if err != nil {
		if errors.Is(err, services.ErrConversationForbidden) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		http.Error(
			w,
			"AI request failed",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
