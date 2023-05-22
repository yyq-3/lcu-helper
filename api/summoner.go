package api

import (
	"encoding/json"
	"lcu-helper/logger"
	"lcu-helper/model"
)

// 召唤师API

// GetCurrentSummonerInfo 获取当前召唤师信息
func GetCurrentSummonerInfo(apiAddr string) *model.UserInfo {
	var user model.UserInfo
	data, err := Get(apiAddr, CURRENT_SUMMONER)
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
