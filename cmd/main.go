package main

import (
	"lcu-helper/internal/lcu"
	"lcu-helper/internal/listener"
	"lcu-helper/internal/os/windows/admin"
	"lcu-helper/pkg/logger"
	"lcu-helper/pkg/tts"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		exit()
	}()
	// 申请管理员权限
	admin.WithAdminRun()
	// 创建窗口
	go func() {
		//ui.Init()
	}()
	// 初始化语音助手
	tts.Init()
	// start process listener
	listener.StartClientListen()
	// 初始化lcu
	lcu.Init()
	// hold main thread
	for {
	}

}

func exit() {
	logger.Info("游戏助手关闭，感谢您的使用~")
	if lcu.Socket.IsConnected {
		lcu.Socket.Close()
	}
	tts.Exit()
	os.Exit(0)
}
