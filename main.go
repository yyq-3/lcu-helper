package main

import (
	"lcu-helper/lcu"
	"lcu-helper/listener"
	"lcu-helper/os/windows/admin"
)

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

func main() {
	// 申请管理员权限
	admin.WithAdminRun()
	// start process listener
	listener.StartClientListen()
	// hold main thread
	lcu.Init()
	for {
	}

}
