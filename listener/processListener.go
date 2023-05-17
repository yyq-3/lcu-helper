package listener

import (
	"lcu-helper/dto"
	"lcu-helper/global"
	"lcu-helper/logger"
	"lcu-helper/util"
	"time"
)

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

func StartClientListen() {
	for {
		if !global.ClientUx.Status {
			handler(global.ClientUx)
		}
		// sleep 1 second
		time.Sleep(time.Duration(time.Second))
	}
}

func handler(client *dto.ClientStatus) {
	if util.ProcessIsRun(client.ProcessName) {
		updateStatus()
		logger.Info("检测到客户端启动")
		logger.Info("开始获取端口和Token")

		logger.Infof("获取到Port: %d, Token: %s")
		logger.Info("开始连接游戏客户端........")
	} else {
		logger.Info("未检测到客户端启动")
	}
}

func updateStatus() {
	global.ClientUx.Lock.Lock()
	defer global.ClientUx.Lock.Unlock()
	global.ClientUx.Status = true
}
