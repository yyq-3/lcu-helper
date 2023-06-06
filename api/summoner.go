package api

// 召唤师API

// GetCurrentSummonerInfo 获取当前召唤师信息
func GetCurrentSummonerInfo(apiAddr string) (data []byte, err error) {
	return get(apiAddr, SummonerCurrent)
}
