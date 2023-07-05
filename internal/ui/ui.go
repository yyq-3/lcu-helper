package ui

import (
	_ "embed"
	"github.com/AllenDang/giu"
	"lcu-helper/internal/os/windows/user32"
)

//go:embed style.css
var cssStyle []byte

var screenWidth, screenHeight int

func Init() {
	screenWidth, screenHeight = user32.GetScreenSize()
	masterWindow := giu.NewMasterWindow("lcu-helper",
		screenWidth, screenHeight,
		giu.MasterWindowFlagsNotResizable|giu.MasterWindowFlagsMaximized|giu.MasterWindowFlagsFloating|giu.MasterWindowFlagsFrameless|giu.MasterWindowFlagsTransparent,
	)
	_ = giu.ParseCSSStyleSheet(cssStyle)
	//masterWindow.SetBgColor(giu.Vec4ToRGBA(imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0}))
	//masterHwnd := user32.FindWindow("GLFW30", "lcu-helper")
	//logger.Infof("hwnd:%d", masterHwnd)
	//user32.SetWindowTransparent(masterHwnd)
	masterWindow.Run(loop)
}

func loop() {
	//giu.SingleWindow().Layout(giu.Label("left top window"))
	giu.Window("left top").Flags(
		giu.WindowFlagsNoResize | giu.WindowFlagsNoMove | giu.WindowFlagsNoCollapse).Layout(
		giu.Label("left top window"))
	giu.Window("right top").Pos(float32(screenWidth-50), 0).Layout(giu.Label("left top window"))
}
