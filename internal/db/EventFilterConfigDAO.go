package db

type EventFilterConfigDO struct {
	Id        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Uri       string `gorm:"unique;column:uri;type:varchar(255)"`
	MatchRule string `gorm:"column:match_rule;type:varchar(10)"`
	Desc      string `gorm:"column:desc;type:varchar(255)"`
}

func (EventFilterConfigDO) TableName() string {
	return "event_filter_config"
}

// QueryAll 查询所有配置
func (e *EventFilterConfigDO) QueryAll() []*EventFilterConfigDO {
	var configs []*EventFilterConfigDO
	db.Select("Uri", "MatchRule").Find(&configs)
	return configs
}
