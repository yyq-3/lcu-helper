package global

import (
	"lcu-helper/model"
)

var JsonEventPrefixLen = len(`[8,"OnJsonApiEvent",`)
var MyGameInfo = &model.UserInfo{}

const (
	GameFlowPhase = "/lol-gameflow/v1/gameflow-phase"
)
