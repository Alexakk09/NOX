package handlers

import (
	"fmt"
	
	"net/http"
	
)

// This function handles requests to "/"
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API is healthy ✅")
}
