package api

import (
	"encoding/json"
	"fmt"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
)

const (
	LOL = "lol"
	TFT = "tft"
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
func (s *Client) GetCurrentGameAllSummoner() *models.SummonerInProcess {
	var res models.SummonerInProcess
	data, err := s.sendGetRequest(SummonerGameUser)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil
	}
	return &res
}

func (s *Client) SummonerGradeQuery(page, size int, puuid string) *models.UserInfo {
	//data, err := s.sendPostRequest(fmt.Sprintf(SummonerRecordByPuuid, puuid, page, size), nil)
	//if err != nil {
	//	return
	//}
	//err = json.Unmarshal(data)
	//if err != nil {
	//	return
	//}
	return nil
}

// GetSummonerGradeByPUuidForLol 通过PUuid查询玩家近十场LOL战绩
func (s *Client) GetSummonerGradeByPUuidForLol(pUuid string) *models.MatchHistoryLol {
	var res models.MatchHistoryLol
	data, err := s.sendGetRequest(fmt.Sprintf(SummonerRecordByPuuid, LOL, pUuid))
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil
	}
	return &res
}

// GetSummonerGradeByPUuidForTft 通过PUuid查询玩家近十场战绩
// pUUid 召唤师puuid
func (s *Client) GetSummonerGradeByPUuidForTft(pUuid string) *models.MatchHistoryTft {
	var tftRes models.MatchHistoryTft
	data, err := s.sendGetRequest(fmt.Sprintf(SummonerRecordByPuuid, TFT, pUuid))
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &tftRes)
	if err != nil {
		return nil
	}
	return &tftRes
}
