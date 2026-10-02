package handlers
import (
	"fmt"
	
	"net/http"

)

// This function handles requests to "/"

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello %s 👋", name)
}