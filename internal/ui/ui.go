package ui

import (
	"github.com/AllenDang/w32"
	"syscall"
	"unsafe"
)

var windowEx w32.HWND

func Init() {
	className := "lcu-helper"
	windowsName := "windows-name"
	classNamePtr, _ := syscall.UTF16PtrFromString(className)
	windowsNamePtr, _ := syscall.UTF16PtrFromString(windowsName)
	screenWidth := int(w32.GetSystemMetrics(w32.SM_CXSCREEN))
	screenHeight := int(w32.GetSystemMetrics(w32.SM_CYSCREEN))
	// 注册窗口类
	wndclassex := w32.WNDCLASSEX{
		Size:      uint32(unsafe.Sizeof(w32.WNDCLASSEX{})),
		Style:     w32.CS_HREDRAW | w32.CS_VREDRAW,
		Instance:  w32.GetModuleHandle(""),
		ClassName: classNamePtr,
		MenuName:  nil,
	}
	w32.RegisterClassEx(&wndclassex)
	windowEx = w32.CreateWindowEx(0, classNamePtr, windowsNamePtr, w32.WS_POPUP,
		0, 0, screenWidth, screenHeight, 0, 0, 0, nil)
	// 窗口透明
	w32.SetWindowLong(windowEx, w32.GWL_EXSTYLE, w32.WS_EX_LAYERED|w32.WS_EX_TRANSPARENT)
	// 鼠标穿透
	w32.SetWindowLong(windowEx, w32.GWL_EXSTYLE, uint32(w32.GetWindowLong(windowEx, w32.GWL_EXSTYLE)|w32.WS_EX_TRANSPARENT))
	// 设置字体
	w32.SendMessage(windowEx, w32.WM_SETFONT, uintptr(createFont()), 1)
	// 窗口显示
	w32.ShowWindow(windowEx, w32.SW_SHOWMAXIMIZED)
	// 窗口置顶
	w32.SetForegroundWindow(windowEx)
	DrawText("写字测试")
	w32.UpdateWindow(windowEx)
	var msg w32.MSG
	for w32.GetMessage(&msg, 0, 0, 0) > 0 {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

// DrawText
// 窗口写字
func DrawText(text string) {
	w32.DrawText(w32.HDC(windowEx), text, -1, &w32.RECT{
		Left:   100,
		Top:    100,
		Right:  500,
		Bottom: 500,
	}, w32.DT_CENTER|w32.DT_VCENTER|w32.DT_SINGLELINE)
}

func createFont() w32.HFONT {
	lf := w32.LOGFONT{
		Height:         40,
		Width:          0,
		Escapement:     0,
		Orientation:    0,
		Weight:         w32.FW_NORMAL,
		Italic:         0,
		Underline:      0,
		StrikeOut:      0,
		CharSet:        w32.ANSI_CHARSET,
		OutPrecision:   w32.OUT_DEFAULT_PRECIS,
		ClipPrecision:  w32.CLIP_DEFAULT_PRECIS,
		Quality:        w32.DEFAULT_QUALITY,
		PitchAndFamily: w32.DEFAULT_PITCH | w32.FF_DONTCARE,
	}
	return w32.CreateFontIndirect(&lf)
}
