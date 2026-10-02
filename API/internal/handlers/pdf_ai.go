package handlers

import (
	"encoding/json"
	"net/http"

	"veerai/internal/ai"
	"veerai/internal/services"
)

type PDFAIHandler struct {
	provider ai.Provider
}

func NewPDFAIHandler(provider ai.Provider) *PDFAIHandler {
	return &PDFAIHandler{
		provider: provider,
	}
}

func (h *PDFAIHandler) Explain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	text, err := getUploadedPDFText(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	explanation, err := services.ExplainPDF(h.provider, text)
	if err != nil {
		http.Error(w, "Failed to generate explanation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"response": explanation,
	})
}

func (h *PDFAIHandler) Summarize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	text, err := getUploadedPDFText(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	summary, err := services.SummarizePDF(h.provider, text)
	if err != nil {
		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"response": summary,
	})
}
