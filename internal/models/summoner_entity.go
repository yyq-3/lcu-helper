package models

import (
	"gorm.io/gorm"
	"time"
)

type SummonerEntity struct {
	Name          string
	Puuid         string
	MatchCount    uint8
	Kills         uint8
	Deaths        uint8
	Assists       uint8
	LastMatchTime time.Time
	gorm.Model
}
