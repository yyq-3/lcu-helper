package route

import (
	"lcu-helper/internal/context"
	"lcu-helper/internal/models"
	"lcu-helper/internal/strategy"
	"strings"
)

// Router 路由器
type Router struct {
	context *context.Context
}

// NewRouter 创建一个新的路由器
func NewRouter() *Router {
	return &Router{
		context: &context.Context{},
	}
}

// Route 根据请求的方法来选择策略
func (r *Router) Route(res *models.WsResponseResult) {
	var serviceStrategy strategy.Strategy

	if strings.HasPrefix(res.Uri, "/lol-chat/v1/conversations/") && strings.HasSuffix(res.Uri, "lol-champ-select.pvp.net/messages") {
		serviceStrategy = &strategy.ChatStrategy{}
	} else {
		return
	}
	//switch method {
	//case "GET":
	//	strategy = &strategy.GetStrategy{}
	//case "POST":
	//	strategy = &strategy.PostStrategy{}
	//default:
	//	return "Unsupported method"
	//}

	r.context.SetStrategy(serviceStrategy)
	r.context.Execute(res)
}
