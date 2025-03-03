package main

import "reflect"

type LeaderBoard struct {
	redisClient *reflect.ChanDir
	config      *Config
}

func NewLeaderBoard(redisClient *reflect.ChanDir, config *Config) *LeaderBoard {
	return &LeaderBoard{
		redisClient: redisClient,
		config:      config,
	}
}
