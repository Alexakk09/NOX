package main

import (
	"log"
	"net/http"
	"os"

	"veerai/internal/ai"
	"veerai/internal/database"
	"veerai/internal/handlers"
	"veerai/internal/routes"
	"veerai/internal/services"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database.Connect()

	baseURL := os.Getenv("LMSTUDIO_BASE_URL")
	model := os.Getenv("LMSTUDIO_MODEL")

	if baseURL == "" {
		log.Fatal("LMSTUDIO_BASE_URL is not set")
	}

	if model == "" {
		log.Fatal("LMSTUDIO_MODEL is not set")
	}

	provider := ai.NewLMStudioProvider(baseURL, model)

	chatService := services.NewChatService(provider)

	chatHandler := handlers.NewChatHandler(chatService)

	pdfAIHandler := handlers.NewPDFAIHandler(provider)

	routes.RegisterRoutes()
	routes.RegisterPDFAIRoutes(pdfAIHandler)
	routes.RegisterChatRoutes(chatHandler)

	log.Println("Server running on :8080")

	err := http.ListenAndServe(
		":8080",
		corsMiddleware(http.DefaultServeMux),
	)
	if err != nil {
		log.Fatal(err)
	}
}