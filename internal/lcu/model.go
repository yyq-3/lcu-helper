package lcu

import (
	"github.com/gorilla/websocket"
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
	"sync"
)

type ClientProcessInfo struct {
	ProcessName   string
	Lock          sync.RWMutex
	Status        bool
	Port          int
	Token         string
	Pid           uint32
	WebSocketAddr string
	ApiAddr       string
	ClientSocket  gowebsocket.Socket
	ClientConn    *websocket.Conn
}

type GameInfo struct {
	MySummonerId    int64
	MySummonerPUuid string
	Team            string
	TeamOne         []string
	TeamTwo         []string
	ChatGroupId     string
	ReConnect       bool
	MapInfo         models.GameFlowSessionData
}

type HuntEvent struct {
	Type           string `json:"type"`
	Timestamp      int64  `json:"timestamp"`
	KillerName     string `json:"killerName"`
	MonsterSubType string `json:"monsterSubType"`
	MonsterType    string `json:"monsterType"`
}

func (game *GameInfo) clear() {
	game.Team = ""
	game.TeamOne = nil
	game.ChatGroupId = ""
	game.TeamTwo = nil
	logger.Info("上次对局信息清理完毕")
}
