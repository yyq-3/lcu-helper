package models

import "time"

type (
	// 聊天组
	Conversation struct {
		GameName           string      `json:"gameName"`
		GameTag            string      `json:"gameTag"`
		Id                 string      `json:"id"`
		InviterId          string      `json:"inviterId"`
		IsMuted            bool        `json:"isMuted"`
		LastMessage        interface{} `json:"lastMessage"`
		Name               string      `json:"name"`
		Password           string      `json:"password"`
		Pid                string      `json:"pid"`
		TargetRegion       string      `json:"targetRegion"`
		Type               string      `json:"type"`
		UnreadMessageCount int         `json:"unreadMessageCount"`
	}
	ConversationMsg struct {
		Body           string    `json:"body"`
		FromId         string    `json:"fromId"`
		FromPid        string    `json:"fromPid"`
		FromSummonerId int64     `json:"fromSummonerId"`
		Id             string    `json:"id"`
		IsHistorical   bool      `json:"isHistorical"`
		Timestamp      time.Time `json:"timestamp"`
		Type           string    `json:"type"`
	}
)
