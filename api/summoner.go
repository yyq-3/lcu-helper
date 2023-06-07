package api

import (
	"encoding/json"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
)

type Summoner struct {
}

// GetCurrentSummonerInfo 获取当前召唤师信息
func (s *Summoner) GetCurrentSummonerInfo(apiAddr string) *models.UserInfo {
	var user models.UserInfo
	data, err := get(apiAddr, SummonerCurrent)
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

func (s *Summoner) Test() {
	logger.Info("test111111111111111111111")
}
