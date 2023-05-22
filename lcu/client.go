package lcu

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/global"
	"lcu-helper/logger"
	"lcu-helper/util"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var ClientUx = &ClientStatus{
	ProcessName: "LeagueClientUx.exe",
	Port:        0,
	Token:       "",
}

var socket gowebsocket.Socket
var lastResponse WsResponseResult

func Init() {
	go initGameFlow()
	//currentSummonerInfo := api.GetCurrentSummonerInfo(ClientUx.ApiAddr)
	//logger.Infof("%v", currentSummonerInfo)
}

func initGameFlow() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket = gowebsocket.New(ClientUx.WebSocketAddr)
	header := http.Header{}
	header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(util.Str2byte(fmt.Sprintf("riot:%s", ClientUx.Token))))
	options := gowebsocket.ConnectionOptions{
		UseSSL: true,
	}
	socket.RequestHeader = header
	socket.ConnectionOptions = options
	socket.OnTextMessage = onTextMessage
	socket.OnConnected = onConnected
	socket.OnDisconnected = onDisconnected
	socket.OnConnectError = onConnectError
	socket.OnPingReceived = onPingReceived
	socket.OnPongReceived = onPongReceived
	socket.Timeout = time.Hour * 12
	logger.Info("开始连接")
	socket.Connect()
	for {
		select {
		case <-interrupt:
			logger.Info("游戏助手关闭，感谢您的使用~")
			socket.Close()
			return
		}
	}
}

func onTextMessage(message string, socket gowebsocket.Socket) {
	if len(message) <= global.JsonEventPrefixLen {
		return
	}
	message = message[global.JsonEventPrefixLen : len(message)-1]
	err := json.Unmarshal(util.Str2byte(message), &lastResponse)
	if err != nil {
		logger.Info("解析客户端消息异常，异常信息 %s", err.Error())
	}
	//logger.Infof("%v", lastResponse)
	switch lastResponse.Uri {
	case global.GameFlowPhase:
		gameFlowPhase(lastResponse.Data)
	}
}

// 游戏状态切换
func gameFlowPhase(data interface{}) {
	status := data.(string)
	if status == "" {
		return
	}
	logger.Infof("游戏状态切换，当前状态：%s", status)
}

func onPongReceived(data string, s gowebsocket.Socket) {
	logger.Infof("收到Pong请求, %s", data)
}

func onPingReceived(data string, s gowebsocket.Socket) {
	logger.Infof("收到Ping请求, %s", data)
}

func onConnectError(err error, socket gowebsocket.Socket) {
	logger.Infof("连接失败, 失败原因: %s", err.Error())
}

func onDisconnected(err error, socket gowebsocket.Socket) {
	logger.Info("连接关闭!!!")
}

func onConnected(socket gowebsocket.Socket) {
	logger.Info("连接到客户端成功!")
	socket.SendText("[5, \"OnJsonApiEvent\"]")
}
