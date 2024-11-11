package strategy

import (
	"encoding/json"
	"fmt"
	"lcu-helper/api"
	"lcu-helper/internal/global"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
	"math"
	"strconv"
	"time"
)

var client = &api.Client{}

type ChatStrategy struct{}

// Handle 处理 chat相关 事件的具体实现
func (g *ChatStrategy) Handle(res *models.WsResponseResult) {
	// create类型为join room，从而可以获取队友五人信息
	var conversationMsgList []models.ConversationMsg
	s, _ := json.Marshal(res.Data)
	_ = json.Unmarshal(s, &conversationMsgList)
	for _, msg := range conversationMsgList {
		if msg.Body == "joined_room" {
			go queryPlayerInfo(msg.FromId)
		}
	}

}

func queryPlayerInfo(puuid string) {
	// 通过puuid查询战绩、信息
	info := client.GetSummonerInfoByPuuid(puuid)
	if info != nil {
		// 判断是否查询过
		global.GameInfo.Lock.Lock()
		if !checkSearch(puuid, global.GameInfo.TeamOne) {
			global.GameInfo.TeamOne = append(global.GameInfo.TeamOne, puuid)
		} else {
			return
		}
		global.GameInfo.Lock.Unlock()
		count := 0
		// 查战绩，可能查不出来，循环几次
		for {
			history := client.GetSummonerGradeByPUuidForLol(puuid)
			if history != nil {
				games := history.Games.Games
				kills, deaths, assists := 0, 0, 0
				score := 0.0
				damageConversionRates := 0.0
				for _, game := range games {
					if count >= global.CalcNum {
						break
					}

					stats := game.Participants[0].Stats
					if stats.Deaths == 0 {
						if stats.TotalDamageTaken == 0 {
							// 本次挂机不算
							continue
						}
						// 避免计算分母为0，1不影响kda计算
						stats.Deaths = 1
					}
					count++

					assists += stats.Assists
					kills += stats.Kills
					deaths += stats.Deaths

					championId := game.Participants[0].ChampionId
					kda := (float64(stats.Kills) + 0.75*float64(stats.Assists)) / float64(stats.Deaths)
					damageConversionRate := 0.0
					if stats.GoldEarned != 0 {
						damageConversionRate = float64(stats.TotalDamageDealtToChampions) / float64(stats.GoldEarned)
					}
					damageConversionRates += damageConversionRate
					logger.Infof("当前英雄：%v，kda：%v，伤害转换率：%v", global.ChampionDataMap[strconv.Itoa(championId)].Name, convert2Num((float64(stats.Kills)+float64(stats.Assists))/float64(stats.Deaths)), convert2Num(damageConversionRate)*100)
					score += (kda * 0.75) + (damageConversionRate * 0.25)
				}
				sendMsg := fmt.Sprintf("玩家%s最近五局比赛得分情况：kda=%.2f,场均伤害转化率=%.2f%%,最终得分=%.2f", info.GameName, float64(kills+assists)/float64(deaths), convert2Num(damageConversionRates/float64(count))*100, score)
				logger.Infof(sendMsg)
				//client.SendMessage2Group(global.GameInfo.ChatGroupId, sendMsg)
				return
			} else {
				// logger.Info("未查询到战绩，等待下次查询")
				time.Sleep(500 * time.Millisecond)
			}
		}

	} else {
		logger.Info("未查询到玩家信息")
	}
}

func convert2Num(num float64) float64 {
	return math.Round(num*100) / 100
}

func checkSearch(puuid string, source []string) bool {
	for _, v := range source {
		if puuid == v {
			return true
		}
	}
	return false
}
