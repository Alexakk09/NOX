package handlers

import (
	"encoding/json"
	"net/http"
	"veerai/internal/models"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	

	info := models.About{
		Name:    "VeerAI",
		Version: "1.0",
		Author:  "Veer",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(info)
}
