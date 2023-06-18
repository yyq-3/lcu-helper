package models

type WsResponseResult struct {
	Data      interface{} `json:"data"`
	EventType string      `json:"eventType"`
	Uri       string      `json:"uri"`
}

type GameFlowSessionData struct {
	GameData struct {
		Queue struct {
			MapId    uint8
			Name     string
			IsRanked bool
		}
	}
}
