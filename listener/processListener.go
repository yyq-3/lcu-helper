package listener

import (
	"fmt"
	"lcu-helper/lcu"
	"lcu-helper/logger"
	"lcu-helper/os/windows/admin"
	"regexp"
	"strconv"
	"time"
)

func StartClientListen() {
	for {
		if !lcu.ClientUx.Status {
			b := handler(lcu.ClientUx)
			if b {
				return
			}
		}
		// sleep 1 second
		time.Sleep(time.Duration(time.Second))
	}
}

func handler(client *lcu.ClientStatus) bool {
	if admin.ProcessIsRun(client.ProcessName) {
		updateStatus()
		logger.Infof("检测到客户端启动,进程PID：%d", lcu.ClientUx.Pid)
		logger.Info("开始获取端口和Token")
		for i := 0; i < 10; i++ {
			if !getPortAndToken() {
				logger.Infof("获取失败,正在进行第%d次重试.....", i+1)
			} else {
				break
			}
		}
		logger.Infof("获取到Port: %d, Token: %s", lcu.ClientUx.Port, lcu.ClientUx.Token)
		logger.Info("开始连接游戏客户端........")
		return true
	} else {
		logger.Info("未检测到客户端启动")
		return false
	}
}

func getPortAndToken() bool {
	cmdline, _ := admin.GetCmdline(lcu.ClientUx.Pid)
	if cmdline == "" {
		return false
	}
	reg := regexp.MustCompile(`--remoting-auth-token=(.+?)" "--app-port=(\d+)"`)
	argArray := reg.FindSubmatch([]byte(cmdline))
	if len(argArray) < 3 {
		return false
	}
	lcu.ClientUx.Token = string(argArray[1])
	port, err := strconv.Atoi(string(argArray[2]))
	if err != nil {
		return false
	}
	lcu.ClientUx.Port = port
	lcu.ClientUx.WebSocketAddr = fmt.Sprintf("wss://127.0.0.1:%d", port)
	lcu.ClientUx.ApiAddr = fmt.Sprintf("https://riot:%s@127.0.0.1:%d", lcu.ClientUx.Token, port)
	logger.Infof("lcuApi %s", lcu.ClientUx.ApiAddr)
	return true
}

func updateStatus() {
	lcu.ClientUx.Lock.Lock()
	defer lcu.ClientUx.Lock.Unlock()
	lcu.ClientUx.Status = true
}
