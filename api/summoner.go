package api

import (
	"encoding/json"
	"fmt"
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
	var res []map[string]interface{}
	data, err := s.sendGetRequest(SummonerGameUser)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	logger.Infof("%v", res)

}

// GetSummonerGradeByPUuid 通过PUuid查询玩家近十场战绩
func (s *Client) GetSummonerGradeByPUuid(pUuid string) {
	var res []map[string]interface{}
	data, err := s.sendGetRequest(fmt.Sprintf(SUMMONER_RECORD_BY_PUUID, pUuid, 0, 10))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	logger.Infof("玩家%s的最近十场战绩为\n%v", res)
}
