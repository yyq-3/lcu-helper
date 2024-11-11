package api

// 召唤师相关
const (
	// SummonerCurrent 获取当前召唤师信息
	// 24年名字修改后不太适用，查询出为原始名字，官方建议替换为puuid查询
	SummonerCurrent = "/lol-summoner/v1/current-summoner"
	// SummonerGameUser 获取本局游戏内所有玩家信息
	SummonerGameUser = "/lol-gameflow/v1/session"
	// SUMMONER_INFO_BY_NAME 通过召唤师名称获取召唤师信息
	SUMMONER_INFO_BY_NAME = "/lol-summoner/v1/summoners?name=%s"
	// SummonerInfoById 通过召唤师Id查询召唤师信息
	SummonerInfoById = "/lol-summoner/v1/summoners/%d"
	// SUMMONER_INFO_BY_PUUID 通过Puuid获取召唤师信息
	SUMMONER_INFO_BY_PUUID = "/lol-summoner/v2/summoners/puuid/%s"
	// SummonerRecordByPuuid 通过puuid查询召唤师战绩
	SummonerRecordByPuuid = "/lol-match-history/v1/products/%s/%s/matches"
	// 从pid获取riot ID （riotId 名字#id）
	SummonerRiotIdByPuuid = "/account/v1/accounts/by-puuid/%s"
	// Returns the player name. 返回玩家名称。 /liveclientdata/activeplayername
	// Get all available data. 获取所有可用数据。 /liveclientdata/allgamedata
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
	// All Players 所有玩家 /liveclientdata/playerlist

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
	// ChampionVersionData
	// 英雄数据
	ChampionVersionData = "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js?ts=%d"
)

// tft
const (
	TftExternal  = "/tft-external/v1/active-game"
	TftMatchInfo = "/tft/v1/matches/{match_id}"
)

// 地图
const (
	Queues = "/lol-game-queues/v1/queues"
)

// Events 事件
const (
	Event_liveclientdata_eventdata = "/liveclientdata/eventdata"
	// 好友状态
	Event_friend_info_change = "/lol-hovercard/v1/friend-info"
	// 任务通知
	Event_missions_service = "/lol-missions/v1/series"
	// 不知道什么
	Event_npe_tutorial = "/lol-npe-tutorial-path/v1/tutorials"
	// 更新好友总数
	Event_friend_counts = "/lol-chat/v1/friend-counts"
	// 好友创建房间
	Event_buddies = "/lol-game-client-chat/v1/buddies"
	// 好友创建房间不知道哪里的
	Event_buddies_v2 = "/lol-game-client-chat/v2/buddies"
	// 删除聊天窗口
	Event_Delete_conversations_active = "/lol-chat/v1/conversations/active"
	// 建议玩家
	Event_suggested_players = "/lol-suggested-players/v1/suggested-players"
	// 地图信息
	Event_challenges = "lol-challenges/v1/challenges/local-player"
)

const (
	Filter_key = "%!p(MISSING)vp.net"
)
