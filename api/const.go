package api

// 召唤师相关
const (
	// SUMMONER_CURRENT 获取当前召唤师信息
	SUMMONER_CURRENT = "/lol-summoner/v1/current-summoner"
	// SUMMONER_GAME_USER 获取本局游戏内所有玩家信息
	SUMMONER_GAME_USER = "/lol-gameflow/v1/session"
	// SUMMONER_INFO_BY_NAME 通过召唤师名称获取召唤师信息
	SUMMONER_INFO_BY_NAME = "/lol-summoner/v1/summoners?name=%s"
	// SUMMONER_INFO_BY_ID 通过召唤师Id查询召唤师信息
	SUMMONER_INFO_BY_ID = "/lol-summoner/v1/summoners/%s"
	// SUMMONER_INFO_BY_PUUID 通过Puuid获取召唤师信息
	SUMMONER_INFO_BY_PUUID = "/lol-summoner/v2/summoners/puuid/%s"
	// SUMMONER_RECORD_BY_PUUID 通过puuid查询召唤师战绩
	SUMMONER_RECORD_BY_PUUID = "/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d"
)

// 聊天相关
const (
	// CHAT_SEND_MESSAGE_TO_CHAT_GROUP 发送消息到聊天组
	CHAT_SEND_MESSAGE_TO_CHAT_GROUP = "/lol-chat/v1/conversations/%s/messages"
	// CHAT_GET_CHAT_GROUP 获取聊天组
	CHAT_GET_CHAT_GROUP = "/lol-champ-select/v1/session"
	// CHAT_SEND_MESSAGE_TO_USERNAME 给指定好友发送消息
	CHAT_SEND_MESSAGE_TO_USERNAME = "/lol-game-client-chat/v1/instant-messages?summonerName=%s&message=%s"
)

// 游戏相关
const (
	// GAME_CURRENT_CAMP 获取当前对局阵营红100 蓝200
	GAME_CURRENT_CAMP = "/lol-champ-select/v1/pin-drop-notification"
)

// 设置相关
const (
	// SETTING_BACKAGEGROUND 修改生涯背景
	SETTING_BACKAGEGROUND = "/lol-summoner/v1/current-summoner/summoner-profile"
	// SETTING_RANK_LEVEL 修改段位
	SETTING_RANK_LEVEL = "/lol-chat/v1/me"
	// SETTING_GAME_CLIENT_STATUS 设置游戏状态
	SETTING_GAME_CLIENT_STATUS = "/lol-chat/v1/me"
)

// 配置相关
const (
	// CONFIG_AUTO_ACCEPT 自动接受对局
	CONFIG_AUTO_ACCEPT = "/lol-matchmaking/v1/ready-check/accept"
	// CONFIG_AUTO_RECONNECT 自动重连
	CONFIG_AUTO_RECONNECT = "/lol-gameflow/v1/reconnect"
	// CONFIG_AUTO_NEXT_GAME 自动开始下一把
	CONFIG_AUTO_NEXT_GAME = "/lol-lobby/v2/play-again"
)
