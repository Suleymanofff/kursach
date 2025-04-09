package handlers

import (
	"encoding/json"
	"net/http"
)

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "success",
		"message": "User profile endpoint",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
