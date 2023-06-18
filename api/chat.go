package api

import (
	"encoding/json"
	"fmt"
	"github.com/thoas/go-funk"
	"lcu-helper/internal/models"
	"lcu-helper/internal/util"
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
	reqBody := map[string]string{}
	reqBody["body"] = msg
	_, err := s.sendPostRequest(fmt.Sprintf(ChatSendMessageToChatGroup, chatGroupId), reqBody)
	//logger.Info(util.Byte2str(data))
	if err != nil {
		logger.Infof("发送消息失败：失败原因：%s", err.Error())
		return
	}
	//logger.Infof("消息发送成功，发送内容：%s", msg)
}

// SendMessage2Game
// 发送消息到游戏进程
func (s *Client) SendMessage2Game(msg string) {
	reqBody := map[string]string{}
	reqBody["body"] = msg
	data, err := s.sendPostRequest(ChatSendMessageToGameProcess, reqBody)
	logger.Info(util.Byte2str(data))
	if err != nil {
		logger.Infof("发送消息失败：失败原因：%s", err.Error())
		return
	}
	logger.Infof("消息发送成功，发送内容：%s", msg)
}

// GetAllSummonerByRoomId
// 选择英雄页面获取友军
func (s *Client) GetAllSummonerByRoomId(romeId string) *[]int64 {
	var res []models.ChatHistory
	var all []int64
	data, err := s.sendGetRequest(fmt.Sprintf(ChatGetAllByRoomId, romeId))
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil
	}
	for _, chatHistory := range res {
		match := chatHistory.Body == "joined_room" && chatHistory.Type == "system" && !funk.ContainsInt64(all, chatHistory.FromSummonerId)
		if match {
			all = append(all, chatHistory.FromSummonerId)
		}
	}
	return &all
}
