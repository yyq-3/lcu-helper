package ui

import (
	g "github.com/AllenDang/giu"
	"github.com/lxn/win"
	"image/color"
	"lcu-helper/internal/os/windows/user32"
	"syscall"
)

var screenWidth, screenHeight int
var hide = false

var masterWindow *g.MasterWindow

func Init() {
	screenWidth, screenHeight = user32.GetScreenSize()
	masterWindow = g.NewMasterWindow("lcu-helper",
		screenWidth, screenHeight-50,
		g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsMaximized|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent,
	)
	masterWindow.SetBgColor(color.RGBA{})
	// 注册热键
	masterWindow.RegisterKeyboardShortcuts(g.WindowShortcut{
		Key:      g.KeyPause,
		Callback: showOrHide,
	})
	// 窗口透明可穿透
	titleNameU16Ptr, _ := syscall.UTF16PtrFromString("lcu-helper")
	handle := win.FindWindow(nil, titleNameU16Ptr)
	win.SetWindowLongPtr(handle, win.GWL_EXSTYLE, win.GetWindowLongPtr(handle, win.GWL_EXSTYLE)|win.WS_EX_LAYERED|win.WS_EX_TRANSPARENT)
	win.SetWindowPos(handle, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_NOACTIVATE)

	masterWindow.Run(loop)
}

func loop() {
	//imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	//g.PushColorWindowBg(color.RGBA{R: 50, G: 50, B: 50})
	//g.PushColorFrameBg(color.RGBA{R: 10, G: 10, B: 10})

	g.Window("left top").Flags(
		g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsNoCollapse).Layout(
		g.Label("left top window"))
	g.Window("right top").Pos(float32(screenWidth-50), 0).Layout(
		g.Label("left top window"),
	)
	g.Child().Layout(g.Label("Child window")).Build()
	//g.PopStyleColor()
	//g.PopStyleColor()
	//imgui.PopStyleVar()
}

func showOrHide() {
	hide = !hide
}
