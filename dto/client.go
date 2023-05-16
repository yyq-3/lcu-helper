package dto

import "sync"

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/
type ClientStatus struct {
	ProcessName string
	Lock        sync.Locker
	Status      bool
}
