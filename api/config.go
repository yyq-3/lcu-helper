package api

import (
	"lcu-helper/internal/util"
	"lcu-helper/pkg/logger"
)

func (s *Client) AutoAccept() bool {
	logger.Info("自动接受")
	_, err := s.sendPostRequest(ConfigAutoAccept, nil)
	return err == nil
}

func (s *Client) AutoNextGame() bool {
	logger.Info("自动下一局游戏")
	_, err := s.sendPostRequest(ConfigAutoNextGame, nil)
	return err == nil
}

func (s *Client) AutoConnect() bool {
	logger.Info("自动连接")
	_, err := s.sendPostRequest(ConfigAutoReconnect, nil)
	return err == nil
}

func (s *Client) ModifyRank() bool {
	logger.Info("修改rank信息")
	body := make(map[string]map[string]string)
	lol := make(map[string]string)
	lol["rankedLeagueQueue"] = "CHALLENGER"
	lol["rankedLeagueTier"] = "II"
	lol["rankedLeagueDivision"] = "RANKED_SOLO_5x5"
	body["lol"] = lol
	res, err := s.sendPutRequest(SettingRankLevel, body)
	if err != nil {
		logger.Infof("Rank modify err : %s", err.Error())
	}
	logger.Infof("修改Rank结果：%s", util.Byte2str(res))
	return err == nil

}
