package controllers


import (
	"encoding/json"
	"net/http"
	"rest/models"
)

func InsertNewReaction(w http.ResponseWriter, r *http.Request) {
	SetupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/json")

	var react models.UserReply
	_ = json.NewDecoder(r.Body).Decode(&react)

	models.NewReactionOnReply(react)

	json.NewEncoder(w).Encode(react)
}

