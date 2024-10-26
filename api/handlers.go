package api

import (
	"encoding/json"
	"net/http"
	"text-classifier/models"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) handleClassify(w http.ResponseWriter, r *http.Request) {
	var input models.TextInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
