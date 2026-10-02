package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"veerai/internal/middleware"
	"veerai/internal/repository"
)

type CreateConversationResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateConversationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID <= 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id, err := repository.CreateConversation(userID)
	if err != nil {
		http.Error(w, "Failed to create conversation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(CreateConversationResponse{
		ID: id,
	})
}

func GetConversationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID <= 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conversations, err := repository.GetConversations(userID)
	if err != nil {
		http.Error(w, "Failed to get conversations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(conversations)
}

func GetConversationMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID <= 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conversationIDText := strings.TrimPrefix(
		r.URL.Path,
		"/api/conversations/",
	)

	conversationIDText = strings.TrimSuffix(
		conversationIDText,
		"/messages",
	)

	conversationID, err := strconv.ParseInt(conversationIDText, 10, 64)
	if err != nil || conversationID <= 0 {
		http.Error(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	conversationUserID, err := repository.GetConversationUserID(
		conversationID,
	)
	if err != nil {
		http.Error(w, "Conversation not found", http.StatusNotFound)
		return
	}

	if conversationUserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	messages, err := repository.GetMessages(conversationID)
	if err != nil {
		http.Error(w, "Failed to get conversation messages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}

func DeleteConversationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok || userID <= 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conversationIDText := strings.TrimPrefix(
		r.URL.Path,
		"/api/conversations/",
	)

	conversationID, err := strconv.ParseInt(conversationIDText, 10, 64)
	if err != nil || conversationID <= 0 {
		http.Error(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	conversationUserID, err := repository.GetConversationUserID(
		conversationID,
	)
	if err != nil {
		http.Error(w, "Conversation not found", http.StatusNotFound)
		return
	}

	if conversationUserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = repository.DeleteConversation(conversationID, userID)
	if err != nil {
		log.Println("Delete conversation error:", err)
		http.Error(w, "Failed to delete conversation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
