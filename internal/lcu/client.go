package lcu

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/api"
	"lcu-helper/internal/global"
	"lcu-helper/internal/models"
	"lcu-helper/internal/util"
	"lcu-helper/pkg/logger"
	"net/http"
	"time"
)

var (
	ClientUx = &ClientProcessInfo{
		ProcessName: "LeagueClientUx.exe",
		Port:        0,
		Token:       "",
	}
	gameInfo        = &GameInfo{}
	currentSummoner = &models.UserInfo{}
)

var (
	apiClient    *api.Client
	Socket       gowebsocket.Socket
	lastResponse models.WsResponseResult
)

func Init() {
	go initGameFlow()
}

func initGameFlow() {
	Socket = gowebsocket.New(ClientUx.WebSocketAddr)
	header := http.Header{}
	header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(util.Str2byte(fmt.Sprintf("riot:%s", ClientUx.Token))))
	options := gowebsocket.ConnectionOptions{
		UseSSL: true,
	}
	Socket.RequestHeader = header
	Socket.ConnectionOptions = options
	Socket.OnTextMessage = onTextMessage
	Socket.OnConnected = onConnected
	Socket.OnDisconnected = onDisconnected
	Socket.OnConnectError = onConnectError
	Socket.OnPingReceived = onPingReceived
	Socket.OnPongReceived = onPongReceived
	Socket.Timeout = time.Hour * 12
	Socket.Connect()
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
	if champSelect == status {
		// 英雄选择页面
		handlerChampSelect()
	} else if matchmaking == status {
		// 排队页面
		handlerMatchmaking()
	} else if home == status {
		handlerHome()
	} else if readyCheck == status {
		// 接受or拒绝页面
		handlerReadyCheck()
	} else if inProgress == status {
		// 英雄选择完毕载入游戏进程
		handlerInProgress()
	} else if none == status {
		// 游戏大厅页面，清除上次游戏记录信息
		gameInfo.clear()
	}
}

func handlerInProgress() {

}

// 匹配到准备/拒绝页面
// 可在该页面配置自动接受对局
func handlerReadyCheck() {
	apiClient.AutoAccept()
	logger.Info("自动接受对局")
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
	// 获取聊天信息组
	// 读取队友信息
	// 计算得分 -》 入库 -》 保存
	// 推送公屏
	// 秒选英雄

	// 自动天赋

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
	go startProxy()
	apiClient = api.Init(ClientUx.ApiAddr)
	for {
		// 连接成功后获取当前用户信息并保存到全局变量里
		currentSummoner = apiClient.GetCurrentSummonerInfo()
		if currentSummoner.SummonerId != 0 {
			logger.Infof("%v", currentSummoner)
			gameInfo.MySummonerPUuid = currentSummoner.Puuid
			break
		}
	}
	socket.SendText("[5, \"OnJsonApiEvent\"]")
}
