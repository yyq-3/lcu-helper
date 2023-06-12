package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"lcu-helper/internal/models"
	"lcu-helper/pkg/logger"
)

var dbPool *gorm.DB

// 隐式执行初始化
func init() {
	db, err := gorm.Open(sqlite.Open("C:\\lcu-helper.db?mode=rw"), &gorm.Config{})
	if err != nil {
		logger.Infof("打开数据库失败,失败原因:%s", err.Error())
		return
	}
	dbPool = db

	err = db.AutoMigrate(&models.SummonerEntity{})
	if err != nil {
		logger.Infof("自动创建表结构失败，失败原因:%s", err.Error())
	}
}

func Test() {
	err := dbPool.AutoMigrate(&models.SummonerEntity{})
	if err != nil {
		logger.Infof("自动创建表结构失败，失败原因:%s", err.Error())
	}
}
