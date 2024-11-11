package context

import (
	"lcu-helper/internal/models"
	"lcu-helper/internal/strategy"
)

// Context 上下文，持有一个策略
type Context struct {
	strategy strategy.Strategy
}

// SetStrategy 设置当前的策略
func (c *Context) SetStrategy(s strategy.Strategy) {
	c.strategy = s
}

// Execute 执行当前策略
func (c *Context) Execute(res *models.WsResponseResult) {
	c.strategy.Handle(res)
}
