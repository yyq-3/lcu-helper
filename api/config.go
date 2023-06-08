package api

import "lcu-helper/pkg/logger"

func (s *Client) AutoAccept() bool {
	logger.Info("自动接受")
	_, err := s.sendPostRequest(ConfigAutoAccept, nil)
	return err == nil
}

func (s *Client) AutoNextGame(body any) bool {
	logger.Info("自动下一局游戏")
	_, err := s.sendPostRequest(ConfigAutoNextGame, body)
	return err == nil
}

func (s *Client) ModifyRank() bool {
	logger.Info("修改rank信息")
	body := make(map[string]string)
	body["rankedLeagueQueue"] = "CHALLENGER"
	body["rankedLeagueTier"] = "I"
	body["rankedLeagueDivision"] = "RANKED_SOLO_5x5"
	_, err := s.sendPostRequest(SettingRankLevel, body)
	if err != nil {
		logger.Infof("Rank modify err : %s", err.Error())
	}
	return err == nil

}
