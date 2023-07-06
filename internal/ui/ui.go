package ui

import (
	"github.com/AllenDang/imgui-go"
	"image/color"
	"lcu-helper/internal/os/windows/user32"
	"lcu-helper/pkg/logger"

	g "github.com/AllenDang/giu"
)

var screenWidth, screenHeight int

func Init() {
	screenWidth, screenHeight = user32.GetScreenSize()
	masterWindow := g.NewMasterWindow("lcu-helper",
		screenWidth, screenHeight-50,
		g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsMaximized|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent,
	)
	masterWindow.SetBgColor(color.RGBA{})
	handle, _ := user32.GetWindowHandle("lcu-helper")
	logger.Infof("handle is :%v", handle)
	logger.Infof("%v", user32.GetWindowLongPtr(handle, user32.GWL_EXSTYLE))
	masterWindow.Run(loop)
}

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorWindowBg(color.RGBA{R: 50, G: 50, B: 50})
	g.PushColorFrameBg(color.RGBA{R: 10, G: 10, B: 10})

	g.Window("left top").Flags(
		g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsNoCollapse).Layout(
		g.Label("left top window"))
	g.Window("right top").Pos(float32(screenWidth-50), 0).Layout(g.Label("left top window"))
	g.PopStyleColor()
	g.PopStyleColor()
	imgui.PopStyleVar()
}
