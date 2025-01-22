package ui

// import (
// 	"fmt"
// 	"log"

// 	"github.com/AllenDang/cimgui-go/imgui"
// )

// type WindowRoom struct {
// 	WindowBase

// 	roomName             string
// 	identify             string
// 	send_input_test_data string
// }

// func (w *WindowRoom) Start() {
// 	w.roomName = ""
// 	w.identify = ""
// 	w.Open = true

// 	w.send_input_test_data = ""
// }

// func (w *WindowRoom) Update() {
// 	if !w.Open {
// 		return
// 	}

// 	imgui.SetNextWindowSizeV(w.Size, imgui.CondFirstUseEver)
// 	imgui.SetNextWindowPosV(w.Pos, imgui.CondFirstUseEver, imgui.Vec2{X: 0, Y: 0})
// 	if imgui.BeginV(w.Title, &w.Open, imgui.WindowFlagsNoCollapse) {
// 		w.Size = imgui.WindowSize()
// 		w.Pos = imgui.WindowPos()
// 		if !w.room.IsConnSuc() {
// 			imgui.Text("Room Name:")
// 			imgui.SetCursorPosX(50)
// 			imgui.InputTextWithHint("##RoomName", "Enter Room Name", &w.roomName, 0, func(data imgui.InputTextCallbackData) int {
// 				return 0
// 			})

// 			imgui.Text("Identify:")
// 			imgui.SetCursorPosX(50)
// 			imgui.InputTextWithHint("##UserId", "", &w.identify, 0, func(data imgui.InputTextCallbackData) int {
// 				return 0
// 			})

// 			imgui.Spacing()
// 			imgui.SetCursorPosX(50)
// 			if imgui.Button("Connect") {
// 				if len(w.roomName) > 0 && len(w.identify) > 0 {
// 					w.room.ConnectBySecret(host, apiKey, apiSecret, w.roomName, w.identify)
// 					w.Title = w.room.GetRoomName()
// 				}
// 			}
// 		} else {
// 			imgui.Text(fmt.Sprintf("Owner:%s", w.room.GetOwner()))
// 			imgui.Text("-------------------Function----------------------")
// 			if imgui.Button("GetAllParticipant") {
// 				log.Println(w.room.GetAllParticipantId())
// 			}
// 			imgui.Text("-------------------Audio Track----------------------")
// 			if w.track.IsOpened() {
// 				imgui.PushIDStr("AudioTrack_Stop")
// 				if imgui.Button("Stop") {
// 					w.track.Close()
// 				}
// 				imgui.PopID()
// 			} else {
// 				imgui.PushIDStr("AudioTrack_Start")
// 				if imgui.Button("Start") {
// 					w.track.Open()
// 					w.room.PublicTrack(w.track)
// 				}
// 				imgui.PopID()
// 			}
// 			imgui.Text("-------------------MicPhone----------------------")
// 			if w.micPhone.CanUse() {
// 				if w.micPhone.Using() {
// 					imgui.PushIDStr("MicPhone_Stop")
// 					if imgui.Button("Stop") {
// 						w.micPhone.Stop()
// 					}
// 					imgui.PopID()
// 				} else {
// 					imgui.PushIDStr("MicPhone_Start")
// 					if imgui.Button("Start") {
// 						w.micPhone.Start()
// 					}
// 					imgui.PopID()
// 				}
// 			} else {
// 				imgui.Text("No MicPhone")
// 			}
// 			imgui.Text("-------------------Speaker----------------------")
// 			if w.speaker.CanUse() {
// 				if w.speaker.Using() {
// 					imgui.PushIDStr("Speaker_Stop")
// 					if imgui.Button("Stop") {
// 						w.speaker.Stop()
// 					}
// 					imgui.PopID()
// 				} else {
// 					imgui.PushIDStr("Speaker_Start")
// 					if imgui.Button("Start") {
// 						w.speaker.Start()
// 					}
// 					imgui.PopID()
// 				}
// 			} else {
// 				imgui.Text("No Speaker")
// 			}
// 			imgui.Text("-------------------Send TestData ----------------------")
// 			imgui.InputTextWithHint("##send_input_test_data", "", &w.send_input_test_data, 0, func(data imgui.InputTextCallbackData) int {
// 				return 0
// 			})
// 			if imgui.Button("Send") {
// 				w.room.PublishData(w.send_input_test_data)
// 				w.send_input_test_data = ""
// 			}
// 		}
// 		imgui.End()
// 	}
// 	if !w.Open {
// 		removeWin(w)
// 	}
// }

// func (w *WindowRoom) Destroy() {
// 	if w.speaker != nil {
// 		w.speaker.Dispose()
// 	}
// 	if w.micPhone != nil {
// 		w.speaker.Dispose()
// 	}
// 	if w.room != nil {
// 		w.room.Disconnect()
// 	}
// }

// func newRoomWindow() *WindowRoom {
// 	win := &WindowRoom{
// 		WindowBase: WindowBase{
// 			Pos:   imgui.Vec2{X: 100, Y: 100},
// 			Size:  imgui.Vec2{X: 500, Y: 300},
// 			Title: "connect coom",
// 			Id:    "Room",
// 		},
// 	}
// 	return win
// }
