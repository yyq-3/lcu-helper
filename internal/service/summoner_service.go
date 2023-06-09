package service

import (
	"lcu-helper/internal/dao"
	"lcu-helper/internal/models"
)

func Save(summonerList []models.SummonerEntity) bool {
	for _, summonerEntity := range summonerList {
		dao.Inster(&summonerEntity)
	}
	return true
}
