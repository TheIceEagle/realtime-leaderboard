package main

import (
	"github.com/google/uuid"
	"time"
)

type GameScore struct {
	UserId    string
	Username  string
	GameType  string
	GameScore string
	Timestamp string
}

func NewGameScore(username, gametype, gamescore, timestamp string) *GameScore {
	return &GameScore{
		UserId:    uuid.New().String(),
		Username:  username,
		GameType:  username,
		GameScore: username,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
