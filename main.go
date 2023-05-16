package lcu_helper

import "lcu-helper/listener"

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

func main() {
	// start process listener
	go listener.StartClientListen()
	// hold main thread
	select {}
}
