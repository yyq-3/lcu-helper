package lcu

import (
	"github.com/gorilla/websocket"
	"github.com/sacOO7/gowebsocket"
	"sync"
)

type WsResponseResult struct {
	Data      interface{} `json:"data"`
	EventType string      `json:"eventType"`
	Uri       string      `json:"uri"`
}

type ClientStatus struct {
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
