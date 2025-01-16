package ui

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/openimsdk/openim-rtc/example/io"
	"github.com/openimsdk/openim-rtc/sdk"
)

const (
	host      = "http://127.0.0.1:7880"
	apiKey    = "APImMfGZ6ECRXbd"
	apiSecret = "3Uxe8kZorbAz59UDKQe23XjM5alE9T0OyCnSNXjQ9r9"
)

type WindowRoom struct {
	WindowBase

	roomName string
	identify string

	room     *sdk.Room
	micPhone *io.MicPhone
	speaker  *io.Speaker
}

func (w *WindowRoom) Start() {
	w.roomName = ""
	w.identify = ""
	w.Open = true

	w.room = sdk.NewRoom()
	w.micPhone = io.NewMicPhone()
	w.speaker = io.NewSpeaker()
}

func (w *WindowRoom) Update() {
	if !w.Open {
		return
	}
	if w.room == nil {
		return
	}
	imgui.SetNextWindowSizeV(w.Size, imgui.CondFirstUseEver)
	imgui.SetNextWindowPosV(w.Pos, imgui.CondFirstUseEver, imgui.Vec2{X: 0, Y: 0})
	if imgui.BeginV(w.Title, &w.Open, imgui.WindowFlagsNoCollapse) {
		w.Size = imgui.WindowSize()
		w.Pos = imgui.WindowPos()
		if !w.room.IsConnSuc() {
			imgui.Text("Room Name:")
			imgui.SetCursorPosX(50)
			imgui.InputTextWithHint("##RoomName", "Enter Room Name", &w.roomName, 0, func(data imgui.InputTextCallbackData) int {
				return 0
			})

			imgui.Text("Identify:")
			imgui.SetCursorPosX(50)
			imgui.InputTextWithHint("##UserId", "", &w.identify, 0, func(data imgui.InputTextCallbackData) int {
				return 0
			})

			imgui.Spacing()
			imgui.SetCursorPosX(50)
			if imgui.Button("Connect") {
				if len(w.roomName) > 0 && len(w.identify) > 0 {
					w.room.ConnectBySecret(host, apiKey, apiSecret, w.roomName, w.identify)
					w.Title = w.room.GetRoomName()
				}
			}
		} else {
			imgui.Text(fmt.Sprintf("Identify:%s", w.room.GetLocalIdentify()))
			imgui.Text("-------------------MicPhone----------------------")
			if w.micPhone.CanUse() {
				if imgui.Button("MicPhone") {
					if w.micPhone.Using() {
						if imgui.Button("Stop") {
							w.micPhone.Start()
						}
					} else {
						if imgui.Button("Start") {
							w.micPhone.Start()
						}
					}

				}
			} else {
				imgui.Text("No MicPhone")
			}
			imgui.Text("-------------------Speaker----------------------")
			if w.speaker.CanUse() {
				if w.speaker.Using() {
					if imgui.Button("Stop") {

					}
				} else {
					if imgui.Button("Start") {

					}
				}
			} else {
				imgui.Text("No Speaker")
			}

			if imgui.Button("Send Data") {
				w.room.PublishData("hello world")
			}
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
			Title: "connect coom",
			Id:    "Room",
		},
	}
	return win
}
