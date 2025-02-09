package ui

import (
	// "math"

	"log"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
	"github.com/openimsdk/openim-rtc/example/core"
)

var micPhoneData []uint32

const MaxMicPhoneDataLength = 500

var speakerData []uint32

const MaxSpeakerDataLength = 500

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
		if imgui.Button("Unpublish Audio Track") {
			core.StopAudioTrack()
		}
	} else {
		if imgui.Button("Public Audio Track") {
			core.OpenAudioTrack()
		}
	}
	imgui.Text("------------------MicPhone------------------------")
	if core.GetMicPhone().Using() {
		if imgui.Button("Close Micphone") {
			err := core.GetMicPhone().Stop()
			if err != nil {
				log.Panicln(err)
			}
		}
		if implot.PlotBeginPlotV("Micphone", imgui.NewVec2(-1, 300), 0) {
			implot.PlotPlotLineU32PtrInt("line", &micPhoneData[0], int32(len(micPhoneData)))
			implot.PlotEndPlot()
		}
	} else {
		if imgui.Button("Open Micphone") {

			err := core.GetMicPhone().Start()
			if err != nil {
				log.Panic(err)
			} else {
				core.GetMicPhone().AddCallBack(func(data []byte, frameCount uint32) {
					sum := uint32(0)
					for i := 0; i < len(data); i += 2 {
						val := uint32(data[i]) | uint32(data[i+1])<<8
						sum += val
					}
					if len(micPhoneData) > MaxMicPhoneDataLength {
						micPhoneData = micPhoneData[:0]
					}
					val := sum / frameCount
					micPhoneData = append(micPhoneData, val)
				})
			}
		}
	}
	imgui.Text("------------------Speaker------------------------")
	if core.GetSpeaker().Using() {
		if imgui.Button("Close Speaker") {
			err := core.GetSpeaker().Stop()
			if err != nil {
				log.Panicln(err)
			}
		}
		if len(speakerData) > 0 {
			if implot.PlotBeginPlotV("Speaker", imgui.NewVec2(-1, 300), 0) {
				implot.PlotPlotLineU32PtrInt("line", &speakerData[0], int32(len(speakerData)))
				implot.PlotEndPlot()
			}
		} else {
			imgui.Text("No Recv data")
		}
	} else {
		if imgui.Button("Open Speaker") {
			err := core.GetSpeaker().Start()
			if err != nil {
				log.Panic(err)
			} else {
				core.GetSpeaker().AddCallBack(func(data []byte) {
					sum := uint32(0)
					for i := 0; i < len(data); i += 2 {
						val := uint32(data[i]) | uint32(data[i+1])<<8
						sum += val
					}
					if len(speakerData) > MaxSpeakerDataLength {
						speakerData = speakerData[:0]
					}
					val := sum / uint32((len(data) / 2))
					speakerData = append(speakerData, val)
				})
			}
		}
	}
	imgui.End()
}

func drawMainWin() {
	drawRoomWin()
	drawAudioTrackWin()
}
