package main

import (
	"fmt"
	"lcu-helper/internal/lcu"
	"lcu-helper/internal/listener"
	"lcu-helper/internal/os/windows/admin"
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
		fmt.Println("程序即将关闭")
		tts.Exit()
		os.Exit(0)
	}()
	// 申请管理员权限
	admin.WithAdminRun()
	tts.Init()
	tts.Speak("12354")
	// start process listener
	listener.StartClientListen()
	//
	// hold main thread
	lcu.Init()
	for {
	}

}
