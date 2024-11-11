package db

import (
	"encoding/json"
	"lcu-helper/internal/util"
)

type EventMsgDO struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Content   string `gorm:"type:longtext;column:content"`
	EventType string `gorm:"type:varchar(255);column:event_type"`
	Uri       string `gorm:"type:varchar(1024);column:uri"`
}

// 自定义表名方法
func (EventMsgDO) TableName() string {
	return "event_msg"
}

func Insert(body interface{}, eventType string, uri string) bool {
	marshal, err := json.Marshal(body)
	if err != nil {
		return false
	}
	eventMsg := EventMsgDO{
		Content:   util.Byte2str(marshal),
		EventType: eventType,
		Uri:       uri,
	}

	db.Create(&eventMsg)
	return true
}
