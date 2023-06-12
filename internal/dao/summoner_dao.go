package dao

import "lcu-helper/internal/models"

func FindbyId(id int64) *models.SummonerEntity {
	var res *models.SummonerEntity
	dbPool.First(&res, id)
	return res
}

func FindById(ids []int64) *[]models.SummonerEntity {
	var res *[]models.SummonerEntity
	dbPool.Find(&res, "id in ?", ids)
	return res
}

func Inster(entity *models.SummonerEntity) {
	dbPool.Create(entity)
}

func Update(entity models.SummonerEntity) {
	dbPool.Model(entity).Update("match_count", entity.MatchCount+1)
}
