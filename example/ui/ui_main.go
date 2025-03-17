package ui

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
)

const AudioFrameCacheLength = 500

var rawAudioData []float64 = make([]float64, AudioFrameCacheLength)
var dspAudioData []float64 = make([]float64, AudioFrameCacheLength)

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

func normalizedRMS(data []byte, sampleCount int) float64 {
	var sumSquares float64
	for i := 0; i < sampleCount; i++ {
		// 将每两个字节转换为一个int16样本值
		sample := int16(binary.LittleEndian.Uint16(data[i*2 : (i+1)*2]))
		// 计算样本值的平方，并累加
		sumSquares += float64(sample * sample)
	}

	// 计算平均值
	meanSquares := sumSquares / float64(sampleCount)
	// 计算平方根，得到RMS值
	rms := math.Sqrt(meanSquares)
	return rms
}

func appendRawAudioFrame(data []byte, frameCount uint32) {
	if len(rawAudioData) > AudioFrameCacheLength {
		rawAudioData = rawAudioData[:0]
	}
	rawAudioData = append(rawAudioData, normalizedRMS(data, int(frameCount)))
}

func appendClearAudioFrame(data []byte, frameCount uint32) {
	if len(dspAudioData) > AudioFrameCacheLength {
		dspAudioData = dspAudioData[:0]
	}
	dspAudioData = append(dspAudioData, normalizedRMS(data, int(frameCount)))
}

func drawAudioWave() {
	imgui.Begin("Audio Track")
	if implot.PlotBeginPlotV("AudioWave", imgui.NewVec2(-1, 500), 0) {
		implot.PlotPushStyleColorVec4(implot.PlotColLine, imgui.NewVec4(1.0, 0, 0, 1.0))
		implot.PlotPlotLinedoublePtrInt("rawaudiowave", &rawAudioData[0], int32(len(rawAudioData)))
		implot.PlotPopStyleColor()
		implot.PlotPushStyleColorVec4(implot.PlotColLine, imgui.NewVec4(0.0, 1.0, 0.0, 1.0))
		implot.PlotPlotLinedoublePtrInt("dspaudiowave", &dspAudioData[0], int32(len(dspAudioData)))
		implot.PlotPopStyleColor()
		implot.PlotEndPlot()
	}
	imgui.End()
}

// func drawAudioRecord() {
// 	if imgui.Begin("AudioRecord") {
// 		if imgui.Button("Start") {
// 			context.MicPhone.Start()
// 		}
// 		if imgui.Button("Stop") {
// 			context.MicPhone.SaveWavFile("out.wav")
// 			context.MicPhone.Stop()
// 		}
// 		imgui.End()
// 	}
// }

func drawMainWin() {
	if context.ConnectState != ConnectSuccess {
		return
	}
	drawLocalParticipant()
	drawRemoteParticipants()
	drawAudioWave()
	// drawAudioRecord()
}
