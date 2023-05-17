package main

import (
	"fmt"
	"github.com/sacOO7/gowebsocket"
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
	go listener.StartClientListen()
	// 读取port和token
	//exec.Command("")
	// 连接socket
	port := uint8(2)
	socketLink := fmt.Sprintf("ws://127.0.0.1:%d", port)
	fmt.Println(socketLink)
	gowebsocket.New(socketLink)

	// hold main thread
	select {}
}
