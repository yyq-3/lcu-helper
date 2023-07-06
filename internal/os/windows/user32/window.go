package user32

import (
	"fmt"
	"syscall"
	"unsafe"
)

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

var (
	user32DLL     = syscall.NewLazyDLL("user32.dll")
	findWindow    = user32DLL.NewProc("FindWindowW")
	getWindowText = user32DLL.NewProc("GetWindowTextW")

	setWindowLongPtr           = user32DLL.NewProc("SetWindowLongPtrW")
	getWindowLongPtr           = user32DLL.NewProc("GetWindowLongPtrW")
	setLayeredWindowAttributes = user32DLL.NewProc("SetLayeredWindowAttributes")
)

func GetWindowHandle(windowTitle string) (syscall.Handle, error) {
	titlePtr, err := syscall.UTF16PtrFromString(windowTitle)
	if err != nil {
		return 0, err
	}

	handle, _, _ := findWindow.Call(0, uintptr(unsafe.Pointer(titlePtr)))
	if handle == 0 {
		return 0, fmt.Errorf("window not found")
	}

	return syscall.Handle(handle), nil
}

const (
	GWL_EXSTYLE       = -20
	WS_EX_TRANSPARENT = 0x00000020
)

func GetWindowLongPtr(handle syscall.Handle, getType int) uintptr {
	r1, _, _ := getWindowLongPtr.Call(uintptr(handle), uintptr(getType))
	return r1
}

func SetWindowLongPtr(handle syscall.Handle) uintptr {
	r1, _, _ := setWindowLongPtr.Call(uintptr(handle))
	return r1
}
