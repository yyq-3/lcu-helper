package strategy

import "lcu-helper/internal/models"

// Strategy 定义一个路由策略接口
type Strategy interface {
	Handle(res *models.WsResponseResult)
}
