package models

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
	SummonerApi   ISummoner
	ConfigApi     IConfig
}
type ISummoner interface {
	GetCurrentSummonerInfo() *UserInfo
	Test()
}

type IConfig interface {
	AutoAccept()
}

//func (c *ClientStatus) GetCurrentSummonerInfo() *UserInfo {
//	var user UserInfo
//	data, err := api.GetCurrentSummonerInfo(c.ApiAddr)
//	if err != nil {
//		logger.Infof("获取召唤师信息失败, %s", err.Error())
//		return nil
//	}
//	err = json.Unmarshal(data, &user)
//	if err != nil {
//		logger.Infof("获取召唤师信息失败, %s", err.Error())
//		return nil
//	}
//	return &user
//}
