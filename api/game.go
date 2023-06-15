package api

import "encoding/json"

// GetLiveClientData 读取游戏数据
func (s *Client) GetLiveClientData() *map[string]interface{} {
	data, err := s.sendGetRequest(GameLiveClientData)
	if err != nil {
		return nil
	}
	var res map[string]interface{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil
	}
	return &res
}
