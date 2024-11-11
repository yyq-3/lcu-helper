package global

import (
	"lcu-helper/internal/models"
)

var JsonEventPrefixLen = len(`[8,"OnJsonApiEvent",`)

const (
	GameFlowPhase   = "/lol-gameflow/v1/gameflow-phase"
	GameFlowSession = "/lol-gameflow/v1/session"

	ChatConversations = `^\/lol-chat\/v1\/conversations\/.*lol-champ-select.pvp.net\/messages`
)

var ChampionData = &models.ChampionVersionData{}
var ChampionDataMap = map[string]models.Hero{}

var GameInfo = &GameInfoAbout{}

const CalcNum = 5
