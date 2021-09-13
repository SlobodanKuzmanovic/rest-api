package controllers

import (
	"encoding/json"
	"net/http"
	"rest/models"
)

func CreateReaction(w http.ResponseWriter, r *http.Request) {
	SetupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/json")

	var reac models.UserQuestion
	_ = json.NewDecoder(r.Body).Decode(&reac)

	models.InsertReaction(reac)

	json.NewEncoder(w).Encode(reac)
}
