package user32

import (
	"fmt"
	"github.com/AllenDang/w32"
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

func FindWindow(className, windowName string) w32.HWND {
	cnPtr, _ := syscall.UTF16PtrFromString(className)
	wnPtr, _ := syscall.UTF16PtrFromString(className)
	return w32.FindWindowW(cnPtr, wnPtr)
}

func SetWindowTransparent(hwnd w32.HWND) {
	oldStyle := w32.GetWindowLongPtr(hwnd, GWL_EXSTYLE)
	newStyle := oldStyle | WS_EX_TRANSPARENT
	w32.SetWindowLongPtr(hwnd, GWL_EXSTYLE, newStyle)
}
