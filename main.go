package lcu_helper

import (
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/listener"
)

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

func main() {
	// 申请管理员权限

	// 连接socket
	port := uint8(2)
	socket := gowebsocket.New(fmt.Sprintf("ws://127.0.0.1:%d", port))
	// start process listener
	go listener.StartClientListen()
	// hold main thread
	select {}
}
