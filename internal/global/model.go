package global

import (
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
	"sync"
)

type GameInfoAbout struct {
	MySummonerId    int64
	MySummonerPUuid string
	Team            string
	TeamOne         []string
	TeamTwo         []string
	ChatGroupId     string
	ReConnect       bool
	MapInfo         models.GameFlowSessionData
	MatchId         string
	Lock            sync.RWMutex
}

func (game *GameInfoAbout) Clear() {
	game.Team = ""
	game.TeamOne = nil
	game.ChatGroupId = ""
	game.TeamTwo = nil
	logger.Info("上次对局信息清理完毕")
}
