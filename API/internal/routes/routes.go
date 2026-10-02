package routes

import (
	"net/http"

	"veerai/internal/handlers"
	"veerai/internal/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	http.HandleFunc("/api/pdf/upload", handlers.UploadPDFHandler)

	http.Handle(
		"/api/conversations",
		middleware.JWTMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					handlers.CreateConversationHandler(w, r)
					return
				}

				if r.Method == http.MethodGet {
					handlers.GetConversationsHandler(w, r)
					return
				}

				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}),
		),
	)

	http.Handle(
		"/api/conversations/",
		middleware.JWTMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					handlers.GetConversationMessagesHandler(w, r)
					return
				}

				if r.Method == http.MethodDelete {
					handlers.DeleteConversationHandler(w, r)
					return
				}

				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}),
		),
	)

	http.Handle(
		"/profile",
		middleware.JWTMiddleware(
			http.HandlerFunc(handlers.ProfileHandler),
		),
	)
}

func RegisterPDFAIRoutes(pdfAIHandler *handlers.PDFAIHandler) {
	http.HandleFunc("/api/pdf/explain", pdfAIHandler.Explain)
	http.HandleFunc("/api/pdf/summarize", pdfAIHandler.Summarize)
}

func RegisterChatRoutes(chatHandler *handlers.ChatHandler) {
	http.Handle(
		"/api/chat",
		middleware.JWTMiddleware(chatHandler),
	)
}
