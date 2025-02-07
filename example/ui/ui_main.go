package ui

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/openimsdk/openim-rtc/example/core"
)

func drawAudioWin() {
	imgui.Begin("Audio")
	if core.HasPublishAudioTrack() {
		if imgui.Button("Close Audio Track") {
			core.StopAudioTrack()
		}
	} else {
		if imgui.Button("Open Audio Track") {
			core.OpenAudioTrack()
		}
	}
	if imgui.Button("Send Hello World") {
		core.SendData("hello world")
	}

	if imgui.Button("Input:Micphone") {

	}

	if imgui.Button("Output:Speaker") {

	}

	imgui.End()
}

func drawMainWin() {
	drawAudioWin()
}
