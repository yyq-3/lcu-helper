package api

import "lcu-helper/internal/models"

var apiAddr string

type Client struct {
}

type Summoner interface {
	GetCurrentSummonerInfo() *models.UserInfo
	GetCurrentGameAllSummoner() *models.SummonerInProcess
	GetSummonerGradeByPUuidForTft(pUuid string) *models.MatchHistoryTft
	GetSummonerGradeByPUuidForLol(pUuid string) *models.MatchHistoryLol
}

type Config interface {
	AutoAccept() bool
	AutoNextGame() bool
	AutoConnect() bool
	ModifyRank() bool
}

type Game interface {
	GetLiveClientData() *map[string]interface{}
}

type Chat interface {
	GetChatGroup() []models.Conversation
	SendMessage2Group(msg string)
}

func Init(addr string) *Client {
	apiAddr = addr
	return &Client{}
}

func (s *Client) sendGetRequest(apiUrl string) ([]byte, error) {
	return get(apiUrl)
}

func (s *Client) sendPostRequest(apiUrl string, body any) ([]byte, error) {
	return post(apiUrl, body)
}

func (s *Client) sendPutRequest(apiUrl string, body any) ([]byte, error) {
	return put(apiUrl, body)
}
