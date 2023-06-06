package main

import (
	"fmt"
	"lcu-helper/lcu"
	"lcu-helper/listener"
	"lcu-helper/os/windows/admin"
	"lcu-helper/pkg"
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
		os.Exit(0)
	}()
	// 申请管理员权限
	admin.WithAdminRun()
	pkg.Speak()
	// start process listener
	listener.StartClientListen()
	//
	// hold main thread
	lcu.Init()
	for {
	}

}
