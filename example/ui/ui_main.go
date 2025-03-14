package ui

import (
	"fmt"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
)

const AudioDataLength = 500

var rawAudioData []uint32 = make([]uint32, AudioDataLength)
var clearAudioData []uint32 = make([]uint32, AudioDataLength)

func drawLocalParticipant() {
	imgui.Begin("LocalParticipayt:" + context.participantName)
	if imgui.Button("Send TestData") {
		context.SendData("Hello From Golang")
	}
	if imgui.Button("Publish Audio Track") {
		context.PublishAudioTrack()
	}

	imgui.End()
}

func drawRemoteParticipants() {
	if imgui.Begin("Remote Participant") {
		rps := context.Room.GetRemoteParticipants()
		if imgui.BeginTable("Remote participants", 3) {
			imgui.TableSetupColumn("Name")
			imgui.TableSetupColumn("Identify")
			imgui.TableSetupColumn("Publication Count")
			imgui.TableHeadersRow()
			// local participant
			imgui.TableNextRow()
			imgui.TableSetColumnIndex(0)
			imgui.Text(context.Room.LocalParticipant.Name())
			imgui.TableSetColumnIndex(1)
			imgui.Text(context.Room.LocalParticipant.Identity())
			imgui.TableSetColumnIndex(2)
			imgui.Text(fmt.Sprintf("%d", len(context.Room.LocalParticipant.TrackPublications())))
			// remote participants
			for _, rp := range rps {
				imgui.TableNextRow()
				imgui.TableSetColumnIndex(0)
				imgui.Text(rp.Name())
				imgui.TableSetColumnIndex(1)
				imgui.Text(rp.Identity())
				imgui.TableSetColumnIndex(2)
				imgui.Text(fmt.Sprintf("%d", len(rp.TrackPublications())))
			}
			imgui.EndTable()
		}
	}
	imgui.End()
}

func appendRawAudioFrame(data []byte, frameCount uint32) {
	sum := uint32(0)
	for i := 0; i < len(data); i += 2 {
		val := uint32(data[i]) | uint32(data[i+1])<<8
		sum += val
	}
	if len(rawAudioData) > AudioDataLength {
		rawAudioData = rawAudioData[:0]
	}
	val := sum / frameCount
	rawAudioData = append(rawAudioData, val)
}

func appendClearAudioFrame(data []byte, frameCount uint32) {
	sum := uint32(0)
	for i := 0; i < len(data); i += 2 {
		val := uint32(data[i]) | uint32(data[i+1])<<8
		sum += val
	}
	if len(clearAudioData) > AudioDataLength {
		clearAudioData = clearAudioData[:0]
	}
	val := sum / frameCount
	clearAudioData = append(clearAudioData, val)
}

func drawAudioWave() {
	imgui.Begin("Audio Track")
	if implot.PlotBeginPlotV("RawAudioWave", imgui.NewVec2(-1, 300), 0) {
		implot.PlotPlotLineU32PtrInt("line", &rawAudioData[0], int32(len(rawAudioData)))
		implot.PlotEndPlot()
	}
	if implot.PlotBeginPlotV("ClerAudioWave", imgui.NewVec2(-1, 600), 0) {
		implot.PlotPlotLineU32PtrInt("line", &clearAudioData[0], int32(len(clearAudioData)))
		implot.PlotEndPlot()
	}
	imgui.End()
}

func drawAudioRecord() {
	if imgui.Begin("AudioRecord") {
		if imgui.Button("Start") {
			context.MicPhone.Start()
		}
		if imgui.Button("Stop") {
			context.MicPhone.SaveWavFile("out.wav")
			context.MicPhone.Stop()
		}
		imgui.End()
	}
}

func drawMainWin() {
	if context.ConnectState != ConnectSuccess {
		return
	}
	drawLocalParticipant()
	drawRemoteParticipants()
	// drawAudioWave()
	drawAudioRecord()
}
