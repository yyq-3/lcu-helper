package api

// 召唤师相关
const (
	// SummonerCurrent 获取当前召唤师信息
	SummonerCurrent = "/lol-summoner/v1/current-summoner"
	// SummonerGameUser 获取本局游戏内所有玩家信息
	SummonerGameUser = "/lol-gameflow/v1/session"
	// SUMMONER_INFO_BY_NAME 通过召唤师名称获取召唤师信息
	SUMMONER_INFO_BY_NAME = "/lol-summoner/v1/summoners?name=%s"
	// SUMMONER_INFO_BY_ID 通过召唤师Id查询召唤师信息
	SUMMONER_INFO_BY_ID = "/lol-summoner/v1/summoners/%s"
	// SUMMONER_INFO_BY_PUUID 通过Puuid获取召唤师信息
	SUMMONER_INFO_BY_PUUID = "/lol-summoner/v2/summoners/puuid/%s"
	// SummonerRecordByPuuid 通过puuid查询召唤师战绩
	SummonerRecordByPuuid = "/lol-match-history/v1/products/%s/%s/matches"
)

// 聊天相关
const (
	// ChatSendMessageToChatGroup 发送消息到聊天组
	ChatSendMessageToChatGroup = "/lol-chat/v1/conversations/%s/messages"
	// ChatSendMessageToGameProcess 发送消息到游戏进程中
	ChatSendMessageToGameProcess = "/lol-chat/v1/conversations/active"
	// ChatGroup 获取选择英雄页面的聊天组
	ChatGroup = "/lol-chat/v1/conversations"
	// ChatGetAllByRoomId 获取选择英雄页面队友信息
	ChatGetAllByRoomId = "/lol-chat/v1/conversations/%s/messages"
	// CHAT_GET_CHAT_GROUP 获取聊天组
	CHAT_GET_CHAT_GROUP = "/lol-champ-select/v1/session"
	// CHAT_SEND_MESSAGE_TO_USERNAME 给指定好友发送消息
	CHAT_SEND_MESSAGE_TO_USERNAME = "/lol-game-client-chat/v1/instant-messages?summonerName=%s&message=%s"
)

// 游戏相关
const (
	// GAME_CURRENT_CAMP 获取当前对局阵营红100 蓝200
	GAME_CURRENT_CAMP = "/lol-champ-select/v1/pin-drop-notification"
	// GAME_PLACE_FRIEND_JOIN_HOME 邀请加入房间
	GAME_PLACE_FRIEND_JOIN_HOME = "/lol-lobby/v2/received-invitations"
	// GAME_CREATE_HOME 创建游戏房间 无限乱斗mapId=19
	GAME_CREATE_HOME = "/lol-lobby/v2/lobby"
	// GameLiveClientData get请求 获取实时数据
	GameLiveClientData = "/liveclientdata/activeplayer"
)

// 设置相关
const (
	// SETTING_BACKAGEGROUND 修改生涯背景
	SETTING_BACKAGEGROUND = "/lol-summoner/v1/current-summoner/summoner-profile"
	// SettingRankLevel 修改段位
	SettingRankLevel = "/lol-chat/v1/me"
	// SETTING_GAME_CLIENT_STATUS 设置游戏状态
	SETTING_GAME_CLIENT_STATUS = "/lol-chat/v1/me"
)

// 配置相关
const (
	// ConfigAutoAccept 自动接受对局
	ConfigAutoAccept = "/lol-matchmaking/v1/ready-check/accept"
	// ConfigAutoReconnect 自动重连
	ConfigAutoReconnect = "/lol-gameflow/v1/reconnect"
	// ConfigAutoNextGame 自动开始下一把
	ConfigAutoNextGame = "/lol-lobby/v2/play-again"
)

// 英雄选择
const (
	CURRENT_CHAMP_SELECT = "/lol-champ-select/v1/current-champion"
)

// tft
const (
	TftExternal = "/tft-external/v1/active-game"
)
