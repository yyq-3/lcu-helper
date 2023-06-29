package ui

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"lcu-helper/internal/os/windows/user32"
)

func Init() {
	screenWidth, screenHeight := user32.GetScreenSize()
	ig := giu.NewMasterWindow("lcu-helper",
		screenWidth, screenHeight,
		,
	)
	ig.Run(loop)
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("Hello world from giu"),
	)
}
