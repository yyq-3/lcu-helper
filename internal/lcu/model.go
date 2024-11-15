package lcu

import (
	"github.com/gorilla/websocket"
	"github.com/sacOO7/gowebsocket"
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

type HuntEvent struct {
	Type           string `json:"type"`
	Timestamp      int64  `json:"timestamp"`
	KillerName     string `json:"killerName"`
	MonsterSubType string `json:"monsterSubType"`
	MonsterType    string `json:"monsterType"`
}
