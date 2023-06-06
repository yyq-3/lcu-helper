package model

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/api"
	"lcu-helper/logger"
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

func (c *ClientStatus) Accept() {
	err := api.AcceptGame(c.ApiAddr)
	if err != nil {
		logger.Infof("error is %s", err.Error())
		return
	}
}

func (c *ClientStatus) GetCurrentSummonerInfo() *UserInfo {
	var user UserInfo
	data, err := api.GetCurrentSummonerInfo(c.ApiAddr)
	if err != nil {
		logger.Infof("获取召唤师信息失败, %s", err.Error())
		return nil
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		logger.Infof("获取召唤师信息失败, %s", err.Error())
		return nil
	}
	return &user
}
