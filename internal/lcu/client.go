package lcu

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/internal/global"
	"lcu-helper/internal/models"
	"lcu-helper/internal/util"
	"lcu-helper/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var ClientUx = &models.ClientStatus{
	ProcessName: "LeagueClientUx.exe",
	Port:        0,
	Token:       "",
}

var socket gowebsocket.Socket
var lastResponse models.WsResponseResult
var currentSummoner = &models.UserInfo{}

func Init() {
	go initGameFlow()
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

func onTextMessage(message string, _ gowebsocket.Socket) {
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
	default:
		//logger.Infof("%v", lastResponse)
	}
}

// 游戏状态切换
func gameFlowPhase(data interface{}) {
	status := data.(string)
	if status == "" {
		return
	}
	logger.Infof("游戏状态切换，当前状态：%s", status)
	if ChampSelect == status {
		// 英雄选择页面
		handlerChampSelect()
	} else if Matchmaking == status {
		// 排队页面
		handlerMatchmaking()
	} else if Home == status {
		handlerHome()
	} else if ReadyCheck == status {
		handlerReadyCheck()
	} else if InProgress == status {
		handlerInProgress()
	}
}

func handlerInProgress() {

}

// 自动接受对局
func handlerReadyCheck() {
	logger.Info("自动接受对局")
	ClientUx.ConfigApi.AutoAccept()
}

func handlerHome() {
	//groupList := api.GetChatGroup(ClientUx.ApiAddr)
	//logger.Infof("%v", *groupList)
}

// 处理排队页面
func handlerMatchmaking() {
	// 获取当前排队房间信息

	logger.Info("开始排队")
}

// 处理英雄选择页面
func handlerChampSelect() {
	//groupList := api.GetChatGroup(ClientUx.ApiAddr)
	//logger.Infof("%v", *groupList)
	logger.Info("进入英雄选择页面")
}

func onPongReceived(data string, _ gowebsocket.Socket) {
	logger.Infof("收到Pong请求, %s", data)
}

func onPingReceived(data string, _ gowebsocket.Socket) {
	logger.Infof("收到Ping请求, %s", data)
}

func onConnectError(err error, socket gowebsocket.Socket) {
	logger.Infof("连接失败, 失败原因: %s, 开始重新连接", err.Error())
	socket.Connect()
}

func onDisconnected(error, gowebsocket.Socket) {
	logger.Info("连接关闭!!!")
}

func onConnected(socket gowebsocket.Socket) {
	logger.Info("连接到客户端成功!")
	go StartProxy()
	for {
		// 连接成功后获取当前用户信息并保存到全局变量里
		currentSummoner = ClientUx.SummonerApi.GetCurrentSummonerInfo()
		if currentSummoner.SummonerId != 0 {
			logger.Infof("%v", currentSummoner)
			break
		}
	}
	socket.SendText("[5, \"OnJsonApiEvent\"]")
}
