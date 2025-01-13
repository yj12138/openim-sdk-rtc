package ui

import (
	"github.com/AllenDang/cimgui-go/imgui"
)

func DrawMainMenu() {
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("Room") {
			if imgui.MenuItemBool("Connect Room") {
				showWindow(newRoomWindow())
			}
			imgui.EndMenu()
		}
		imgui.EndMainMenuBar()
	}
}