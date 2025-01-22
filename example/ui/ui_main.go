package ui

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/openimsdk/openim-rtc/example/core"
)

func drawMainWin() {
	imgui.Begin("Audio")
	if imgui.Button("Publish Audio Track") {
		core.OpenAudioTrack()
	}
	if imgui.Button("UnPublish Audio Track") {
		core.StopAudioTrack()
	}
	if imgui.Button("Send Test Data") {
		core.SendData("hello world")
	}
	if imgui.Button("Open Micphone") {
		core.OpenMicPhone()
	}
	if imgui.Button("Stop Micphone") {
		core.StopMicPhone()
	}
	if imgui.Button("Open Speaker") {
		core.OpenSpeaker()
	}
	if imgui.Button("Stop Speaker") {
		core.StopSpeakder()
	}
	imgui.End()
}
