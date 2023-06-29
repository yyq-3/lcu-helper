package ui

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"lcu-helper/internal/os/windows/user32"
)

func Init() {
	screenWidth, screenHeight := user32.GetScreenSize()
	masterWindow := giu.NewMasterWindow("lcu-helper",
		screenWidth, screenHeight,
		giu.MasterWindowFlagsNotResizable|giu.MasterWindowFlagsMaximized|giu.MasterWindowFlagsFloating|giu.MasterWindowFlagsFrameless|giu.MasterWindowFlagsTransparent,
	)
	masterWindow.SetBgColor(giu.Vec4ToRGBA(imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0}))
	masterWindow.Run(loop)
}

func loop() {
	//giu.SingleWindow().Layout(giu.Label("left top window"))
	giu.Window("left top").Flags(
		giu.WindowFlagsNoResize | giu.WindowFlagsNoMove | giu.WindowFlagsNoCollapse).Layout(
		giu.Label("left top window"))
	giu.Window("right top").Pos(-50, 0).Layout(giu.Label("left top window"))
}
