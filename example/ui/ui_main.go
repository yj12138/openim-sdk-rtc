package ui

import (
	"log"

	"github.com/AllenDang/cimgui-go/imgui"
)

func drawMainWin() {
	// imgui.Begin("mainwindow")
	if imgui.Button("Button") {
		log.Println("click")
	}
	// imgui.End()
}
