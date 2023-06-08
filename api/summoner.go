package api

import (
	"encoding/json"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
)

// GetCurrentSummonerInfo 获取当前召唤师信息
func (s *Client) GetCurrentSummonerInfo() *models.UserInfo {
	var user models.UserInfo
	data, err := s.sendGetRequest(SummonerCurrent)
	if err != nil {
		logger.Infof("获取召唤师信息失败, %s", err.Error())
		return nil
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		logger.Infof("获取召唤师信息失败, %s", err.Error())
		return nil
	}
	return &user
}

// GetCurrentGameAllSummoner 获取本局游戏全部召唤师
func (s *Client) GetCurrentGameAllSummoner() {
	//data, err := s.sendGetRequest(SummonerGameUser)
	//if err != nil {
	//	return
	//}

}
