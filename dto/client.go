package dto

import "sync"

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/
type ClientStatus struct {
	ProcessName string
	Lock        sync.RWMutex
	Status      bool
	Port        int
	Token       string
}
