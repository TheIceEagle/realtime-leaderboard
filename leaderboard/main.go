package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	cfg := LoadConfig()

	client := NewClient(cfg)

	defer client.Close()

	//h := &Handler{payment: paymentService}

}

type Handler struct {
	addScore *LeaderBoard
}

type ScoreRequest struct {
	UserID    string `json:"user_id"`
	GameType  string `json:"game_type"`
	GameScore string `json:"game_score"`
}

type ScoreResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h *Handler) AddScore(w http.ResponseWriter, r *http.Request) {
	var request ScoreRequest

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add score to leaderboard
	err := h.addScore.AddScore(request.UserID, request.GameType, request.GameScore)

}
