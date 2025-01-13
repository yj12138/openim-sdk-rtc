package ui

import (
	"log"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/openimsdk/openim-rtc/sdk"
)

type WindowRoom struct {
	WindowBase

	roomName string
	userId   string

	room *sdk.Room
}

func (w *WindowRoom) Start() {
	w.roomName = ""
	w.userId = ""
	w.Open = true
}

func (w *WindowRoom) Update() {
	if !w.Open {
		return
	}

	imgui.SetNextWindowSizeV(w.Size, imgui.CondFirstUseEver)
	imgui.SetNextWindowPosV(w.Pos, imgui.CondFirstUseEver, imgui.Vec2{X: 0, Y: 0})
	if imgui.BeginV(w.Title, &w.Open, imgui.WindowFlagsNoCollapse) {
		w.Size = imgui.WindowSize()
		w.Pos = imgui.WindowPos()

		imgui.Text("Room Name:")
		imgui.SetCursorPosX(50)
		imgui.InputTextWithHint("##RoomName", "Enter Room Name", &w.roomName, 0, func(data imgui.InputTextCallbackData) int {
			return 0
		})

		imgui.Text("User Id:")
		imgui.SetCursorPosX(50)
		imgui.InputTextWithHint("##UserId", "", &w.userId, 0, func(data imgui.InputTextCallbackData) int {
			return 0
		})

		imgui.Spacing()
		imgui.SetCursorPosX(50)
		if imgui.Button("Connect") {
			log.Println("Try Connect ...")
			// w.room = sdk.NewRoom()
		}

		imgui.End()
	}

}

func (w *WindowRoom) Destroy() {

}

func newRoomWindow() *WindowRoom {
	win := &WindowRoom{
		WindowBase: WindowBase{
			Pos:   imgui.Vec2{X: 100, Y: 100},
			Size:  imgui.Vec2{X: 500, Y: 300},
			Title: "Room",
			Id:    "Room",
		},
	}
	return win
}
