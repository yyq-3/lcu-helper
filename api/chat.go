package api

import (
	"encoding/json"
	"fmt"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
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

func (s *Client) SendMessage2Group(chatGroupId, msg string) {
	req := map[string]string{}
	req["body"] = msg
	body, err := json.Marshal(req)
	if err != nil {
		logger.Infof("发送消息失败：失败原因：%s", err.Error())
		return
	}
	_, err = s.sendPostRequest(fmt.Sprintf(ChatSendMessageToChatGroup, chatGroupId), body)
	if err != nil {
		logger.Infof("发送消息失败：失败原因：%s", err.Error())
		return
	}
	logger.Infof("消息发送成功，发送内容：%s", msg)
}
