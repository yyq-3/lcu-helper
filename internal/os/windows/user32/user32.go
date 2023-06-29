package user32

import "syscall"

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

const (
	SM_CSSCREEN = 0
	SC_CYSCREEN = 1
)

func GetScreenSize() (int, int) {
	user32 := syscall.MustLoadDLL("user32.dll")
	proc := user32.MustFindProc("GetSystemMetrics")
	width, _, _ := proc.Call(uintptr(SM_CSSCREEN))
	height, _, _ := proc.Call(uintptr(SC_CYSCREEN))
	return int(width), int(height)
}
