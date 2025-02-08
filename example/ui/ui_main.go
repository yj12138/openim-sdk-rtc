package ui

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
	"github.com/openimsdk/openim-rtc/example/core"
)

var waveValues [1024]uint32
var index = 0

func init() {
}

func drawRoomWin() {
	imgui.Begin("Room")
	if imgui.Button("Send Hello World") {
		core.SendData("hello world")
	}
	imgui.End()
}

func drawAudioTrackWin() {
	imgui.Begin("Audio Track")
	if core.HasPublishAudioTrack() {
		if imgui.Button("Close Audio Track") {
			core.StopAudioTrack()
		}
		imgui.Text("------------------MicPhone------------------------")
		if core.GetMicPhone().Using() {
			if imgui.Button("Close Micphone") {
				core.GetMicPhone().Stop()
			}
			if implot.PlotBeginPlotV("Micphone", imgui.NewVec2(-1, 300), 0) {
				implot.PlotPlotLineU32PtrInt("line", &waveValues[0], int32(len(waveValues)))
				implot.PlotEndPlot()
			}
		} else {
			if imgui.Button("Open Micphone") {
				core.GetMicPhone().Start()
				core.GetMicPhone().AddCallBack(func(data []byte, frameCount uint32) {
					for i := 0; i < len(data); i++ {
						index = (index + 1) % len(waveValues)
						if index >= 0 && index < len(waveValues) {
							waveValues[index] = uint32(data[i])
						}
					}
				})
			}
		}
		imgui.Text("------------------Speaker------------------------")
		if core.GetSpeaker().Using() {
			if imgui.Button("Close Speaker") {
				core.GetSpeaker().Stop()
			}
		} else {
			if imgui.Button("Open Speaker") {
				core.GetSpeaker().Start()
			}
		}
	} else {
		if imgui.Button("Open Audio Track") {
			core.OpenAudioTrack()
		}
	}
	imgui.End()
}

func drawMainWin() {
	drawRoomWin()
	drawAudioTrackWin()
}
