package api

import (
	"encoding/json"
	"lcu-helper/internal/models"
)

func (s *Client) GetChatGroup() []models.Conversation {
	var conversationList []models.Conversation
	data, err := s.sendGetRequest(ChatGroup)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &conversationList)
	if err != nil {
		return nil
	}
	return conversationList
}
