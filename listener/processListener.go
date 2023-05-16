package listener

import (
	"lcu-helper/dto"
	"time"
)

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

func StartClientListen() {
	clientStatus := dto.ClientStatus{}
	for {
		handler(clientStatus)
		// sleep 1 second
		time.Sleep(time.Duration(time.Second))
	}
}

func handler(client dto.ClientStatus) {

}

func updateStatus(client dto.ClientStatus) {
	client.Lock.Lock()
	defer client.Lock.Unlock()
	client.Status = true
}
