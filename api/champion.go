package api

import (
	"encoding/json"
	"fmt"
	"lcu-helper/internal/global"
	"time"
)

// 初始化英雄数据
func init() {
	// 计算ts
	ts := time.Now().Unix() / 600
	reqUrl := fmt.Sprintf(ChampionVersionData, ts)
	data, err := get(reqUrl)
	if err != nil {
		fmt.Printf("英雄数据初始化失败,失败原因: %s\n", err.Error())
		return
	}
	err = json.Unmarshal(data, global.ChampionData)
	if err != nil {
		fmt.Printf("英雄数据初始化失败,失败原因: %s\n", err.Error())
		return
	}
	fmt.Printf("当前版本[%s],英雄数据初始化完成\n", global.ChampionData.Version)
}
