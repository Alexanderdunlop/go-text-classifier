package api

import (
	"encoding/json"
	"net/http"
	"text-classifier/models"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) handleClassify(w http.ResponseWriter, r *http.Request) {
	var input models.TextInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errorResponse := ErrorResponse{
			Error:   "Bad Request",
			Code:    http.StatusBadRequest,
			Message: "Invalid JSON payload",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if input.Text == "" {
		errorResponse := ErrorResponse{
			Error:   "Bad Request",
			Code:    http.StatusBadRequest,
			Message: "Text field cannot be empty",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// TODO: Add classification logic
	result := models.TextResult{
		Text:     input.Text,
		Category: "placeholder",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
