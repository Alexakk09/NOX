package handlers

import (
	"fmt"
	
	"net/http"
	
)

// This function handles requests to "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to VeerAI API 🚀")
}
