package lcu

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"github.com/thoas/go-funk"
	"lcu-helper/api"
	"lcu-helper/internal/db"
	"lcu-helper/internal/global"
	"lcu-helper/internal/models"
	"lcu-helper/internal/route"
	"lcu-helper/internal/util"
	"lcu-helper/pkg/logger"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ClientUx = &ClientProcessInfo{
		ProcessName: "LeagueClientUx.exe",
		Port:        0,
		Token:       "",
	}
	currentSummoner = &models.SummonerInfo{}
)

var (
	apiClient    *api.Client
	Socket       gowebsocket.Socket
	lastResponse *models.WsResponseResult
)

var eventFilterConfigs []*db.EventFilterConfigDO

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
	if filterEvent2(lastResponse) {
		return
	}
	r := route.NewRouter()
	r.Route(lastResponse)
	switch lastResponse.Uri {
	case global.GameFlowPhase:
		gameFlowPhase(lastResponse.Data)
	case global.GameFlowSession:
		gameFlowSession(lastResponse.Data)
	default:
		db.Insert(lastResponse.Data, lastResponse.EventType, lastResponse.Uri)
		// logger.Infof("default未知客戶端消息：%v", lastResponse)
	}
}

func filterEvent2(response *models.WsResponseResult) bool {
	eventType := response.EventType
	if eventType == "Delete" {
		return true
	}
	uri := response.Uri
	for _, config := range eventFilterConfigs {
		if config.MatchRule == "left" {
			if strings.HasPrefix(uri, config.Uri) {
				return true
			}
		} else if config.MatchRule == "RegEx" {
			result, err := regexp.MatchString(config.Uri, uri)
			if err == nil && result {
				return result
			}
		} else {
			logger.Info("规则错误没有对应的rule")
		}
	}
	return false
}

// 游戏房间信息变化
func gameFlowSession(data interface{}) {
	var res models.GameFlowSessionData
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(marshal, &res)
	if err != nil {
		return
	}
	global.GameInfo.MapInfo = res
	logger.Infof("切换模式，当前模式[%s]", res.GameData.Queue.Name)
}

// 游戏状态切换
func gameFlowPhase(data interface{}) {
	status := data.(string)
	if status == "" {
		return
	}
	logger.Infof("游戏状态切换，当前状态：%s", status)
	if champSelect == status {
		if global.GameInfo.MapInfo.GameData.Queue.MapId == yunDing {
			return
		}
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
		global.GameInfo.Clear()
	} else if reconnect == status {
		apiClient.AutoConnect()
	}
}

func handlerInProgress() {
	go readAttackSpeed()
	var summonerInProcess *models.SummonerInProcess
	// if len(res) not eq 10, every 500ms call
	for {
		summonerInProcess = apiClient.GetCurrentGameAllSummoner()
		if summonerInProcess != nil {
			logger.Info("获取到本局所有玩家信息")
			break
		}
		logger.Info("即将进行下一次查询")
		time.Sleep(time.Millisecond * 500)
	}
	gameMode := summonerInProcess.Map.GameMode
	switch gameMode {
	case "TFT":
		handlerTft(summonerInProcess)
	default:
		handlerLol(summonerInProcess)
	}

}

func readAttackSpeed() {

}

// 处理非云顶游戏
func handlerLol(s *models.SummonerInProcess) {
	allPuuid := make([]string, 5)
	//	msgTemplate := `敌方玩家【%s】最近战绩
	//使用英雄【%s】 战绩【%d/%d/%d】 经济转化率【功能暂未开发】
	//使用英雄【%s】 战绩【%d/%d/%d】 经济转化率【功能暂未开发】
	//使用英雄【%s】 战绩【%d/%d/%d】 经济转化率【功能暂未开发】
	//使用英雄【%s】 战绩【%d/%d/%d】 经济转化率【功能暂未开发】
	//使用英雄【%s】 战绩【%d/%d/%d】 经济转化率【功能暂未开发】`
	if funk.ContainsString(global.GameInfo.TeamOne, s.GameData.TeamOne[0].Puuid) {
		for _, team := range s.GameData.TeamTwo {
			allPuuid = append(allPuuid, team.Puuid)
		}
	} else {
		for _, team := range s.GameData.TeamOne {
			allPuuid = append(allPuuid, team.Puuid)
		}
	}

	for _, puuid := range allPuuid {
		go func(puuid string) {
			var lol *models.MatchHistoryLol
			for {
				lol = apiClient.GetSummonerGradeByPUuidForLol(puuid)
				if lol != nil {
					break
				}
				time.Sleep(time.Second)
			}
			// 发送到游戏
			//apiClient.SendMessage2Game(msgTemplate)
			//tts.Speak(msgTemplate)
		}(puuid)
	}
}

func handlerTft(s *models.SummonerInProcess) {
	for _, teams := range s.GameData.TeamOne {
		if teams.Puuid == global.GameInfo.MySummonerPUuid {
			continue
		}
		// 利用携程查询每个人的对局信息
		go func(puuid, name string) {
			for {
				tftGrade := apiClient.GetSummonerGradeByPUuidForTft(puuid)
				if tftGrade != nil {
					// 循环每局棋子,只处理最近五场的
					for i, game := range tftGrade.Games {
						if i >= 5 {
							break
						}
						for _, p := range game.Json.Participants {
							if p.Puuid == puuid {
								speakBody := fmt.Sprintf("查询到召唤师%s最近阵容", name)
								logger.Info(speakBody)
								//tts.Speak(speakBody)
								break
							}
						}
					}
					break
				}
				time.Sleep(time.Second)
			}
		}(teams.Puuid, teams.SummonerName)
	}
}

// 匹配到准备/拒绝页面
// 可在该页面配置自动接受对局
func handlerReadyCheck() {
	if apiClient.AutoAccept() {
		logger.Info("自动接受对局")
	}
}

func handlerHome() {
	//groupList := api.GetChatGroup(ClientUx.ApiAddr)
	//logger.Infof("%v", *groupList)
}

// 处理排队页面
func handlerMatchmaking() {
	logger.Info("开始排队")
}

// 处理英雄选择页面
func handlerChampSelect() {
	logger.Info("进入英雄选择页面")
	// 获取聊天信息组
	for {
		groupList := apiClient.GetChatGroup()
		if groupList != nil && len(groupList) > 0 {
			logger.Infof("获取到聊天组ID：%s", groupList[0].Id)
			// save id to global
			global.GameInfo.ChatGroupId = groupList[0].Id
			break
		}
		// every 500ms call
		time.Sleep(time.Millisecond * 500)
	}
	// 读取队友信息
	// readTeamSummonerHistory()
	// 计算得分 -》 入库 -》 保存

	// 秒选英雄

	// 自动天赋

}

// 读取团队召唤师对局历史
func readTeamSummonerHistory() {
	for i := 0; i < 5; i++ {
		allUser := apiClient.GetAllSummonerByRoomId(global.GameInfo.ChatGroupId)
		if allUser != nil && len(*allUser) > 0 {
			for _, id := range *allUser {
				//if id == gameInfo.MySummonerId {
				//	continue
				//}
				go func(id int64) {
					summonerInfo := apiClient.GetSummonerInfoById(id)
					if summonerInfo == nil {
						logger.Infof("玩家%d信息获取失败", id)
						return
					}
					puuid := summonerInfo.Puuid
					// 查询战绩 查询不到重试五次
					for i := 0; i < 5; i++ {
						lolHistory := apiClient.GetSummonerGradeByPUuidForLol(puuid)
						if lolHistory != nil {
							analyseLolHistory(lolHistory, summonerInfo)
							break
						}
						logger.Info("正在进行战绩查询重试")
						time.Sleep(time.Second)
					}
				}(id)
			}
			break
		}
		logger.Info("即将进行获取聊天记录重试")
		time.Sleep(time.Millisecond * 500)
	}
}

// 分析玩家历史记录并发送公屏
func analyseLolHistory(history *models.MatchHistoryLol, summonerInfo *models.SummonerInfo) {
	msgTemplate := "最近使用[%s],战绩【%d/%d/%d】,输出%d,补兵%d,经济%d,经济转换率%s%%"
	res := make([]string, 5)
	message := ""
	if summonerInfo.NameChangeFlag {
		message = fmt.Sprintf("玩家【%s】%d级,改过名字,原名称【%s】\r\n", summonerInfo.DisplayName, summonerInfo.SummonerLevel, summonerInfo.InternalName)
	} else {
		message = fmt.Sprintf("玩家【%s】%d级\r\n", summonerInfo.DisplayName, summonerInfo.SummonerLevel)
	}
	if len(history.Games.Games) == 0 {
		message += "近期暂无对局"
		goto sendMsg
	}
	// 取最近五场战绩进行分析
	for i := 0; i < 5; i++ {
		if i >= len(history.Games.Games) {
			break
		}
		game := history.Games.Games[i]
		participant := game.Participants[0]
		championId := participant.ChampionId
		kills := participant.Stats.Kills
		assists := participant.Stats.Assists
		deaths := participant.Stats.Deaths
		totalDamageDealtToChampions := participant.Stats.TotalDamageDealtToChampions
		totalMinionsKilled := participant.Stats.TotalMinionsKilled
		goldEarned := participant.Stats.GoldEarned
		res = append(res,
			fmt.Sprintf(msgTemplate,
				strconv.Itoa(championId),
				//global.ChampionData.Hero[championId+1].Name+"-"+global.ChampionData.Hero[championId+1].Title,
				kills, assists, deaths,
				totalDamageDealtToChampions, totalMinionsKilled, goldEarned,
				fmt.Sprintf("%.3f", float64(totalDamageDealtToChampions*100)/float64(goldEarned))),
		)
	}
sendMsg:
	for _, msg := range res {
		message += msg
		message += "\r\n"
	}
	// 推送公屏
	//apiClient.SendMessage2Group(gameInfo.ChatGroupId, message)
}

func onConnected(socket gowebsocket.Socket) {
	logger.Info("连接到客户端成功!")
	configDO := db.EventFilterConfigDO{}
	eventFilterConfigs = configDO.QueryAll()
	logger.Infof("获取%d条规则", len(eventFilterConfigs))
	go startProxy()
	apiClient = api.Init(ClientUx.ApiAddr)
	for {
		// 连接成功后获取当前用户信息并保存到全局变量里
		currentSummoner = apiClient.GetCurrentSummonerInfo()
		if currentSummoner.SummonerId != 0 {
			logger.Infof("%v", currentSummoner)
			global.GameInfo.MySummonerPUuid = currentSummoner.Puuid
			global.GameInfo.MySummonerId = currentSummoner.SummonerId
			break
		}
	}
	// 获取所有地图信息
	// 修改rank
	apiClient.ModifyRank()
	socket.SendText("[5, \"OnJsonApiEvent\"]")
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
