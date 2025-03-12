package ui

import (
	"fmt"
	"log"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	"github.com/openimsdk/openim-rtc/sdk"
)

// var micPhoneData []uint32

// const MaxMicPhoneDataLength = 500

// var speakerData []uint32

// const MaxSpeakerDataLength = 500

func drawLocalParticipant() {
	imgui.Begin("LocalParticipayt:" + context.participantName)
	if imgui.Button("Send TestData") {
		go func() {
			err := context.LocalParticipant.SendData("Test", []byte("golang hello"), true, []string{})
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
	if imgui.Button("Publish Audio Track") {
		go func() {
			audioSource := sdk.NewAudioSource(audio_frame.AudioSourceType_AUDIO_SOURCE_NATIVE, context.MicPhone.SampleRate, context.MicPhone.Channels)
			track := sdk.NewAudioTrack("micphone audio track", audioSource)
			context.PublishAudioTrack(track)
			context.MicPhone.Start()
			context.MicPhone.AddCallBack(func(data []byte, frameCount uint32) {
				err := audioSource.CaptureFrame(data)
				if err != nil {
					log.Println(err.Error())
				}
			})
		}()
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

func drawAudioTrack() {
	// imgui.Begin("Audio Track")
	// if core.HasPublishAudioTrack() {
	// 	if imgui.Button("Unpublish Audio Track") {
	// 		core.StopAudioTrack()
	// 	}
	// } else {
	// 	if imgui.Button("Public Audio Track") {
	// 		core.OpenAudioTrack()
	// 	}
	// }
	// imgui.Text("------------------MicPhone------------------------")
	// if core.GetMicPhone().Using() {
	// 	if imgui.Button("Close Micphone") {
	// 		err := core.GetMicPhone().Stop()
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}
	// 	}
	// 	if implot.PlotBeginPlotV("Micphone", imgui.NewVec2(-1, 300), 0) {
	// 		implot.PlotPlotLineU32PtrInt("line", &micPhoneData[0], int32(len(micPhoneData)))
	// 		implot.PlotEndPlot()
	// 	}
	// } else {
	// 	if imgui.Button("Open Micphone") {

	// 		err := core.GetMicPhone().Start()
	// 		if err != nil {
	// 			log.Panic(err)
	// 		} else {
	// 			core.GetMicPhone().AddCallBack(func(data []byte, frameCount uint32) {
	// 				sum := uint32(0)
	// 				for i := 0; i < len(data); i += 2 {
	// 					val := uint32(data[i]) | uint32(data[i+1])<<8
	// 					sum += val
	// 				}
	// 				if len(micPhoneData) > MaxMicPhoneDataLength {
	// 					micPhoneData = micPhoneData[:0]
	// 				}
	// 				val := sum / frameCount
	// 				micPhoneData = append(micPhoneData, val)
	// 			})
	// 		}
	// 	}
	// }
	// imgui.Text("------------------Speaker------------------------")
	// if core.GetSpeaker().Using() {
	// 	if imgui.Button("Close Speaker") {
	// 		err := core.GetSpeaker().Stop()
	// 		if err != nil {
	// 			log.Panicln(err)
	// 		}
	// 	}
	// 	if len(speakerData) > 0 {
	// 		if implot.PlotBeginPlotV("Speaker", imgui.NewVec2(-1, 300), 0) {
	// 			implot.PlotPlotLineU32PtrInt("line", &speakerData[0], int32(len(speakerData)))
	// 			implot.PlotEndPlot()
	// 		}
	// 	} else {
	// 		imgui.Text("No Recv data")
	// 	}
	// } else {
	// 	if imgui.Button("Open Speaker") {
	// 		err := core.GetSpeaker().Start()
	// 		if err != nil {
	// 			log.Panic(err)
	// 		} else {
	// 			core.GetSpeaker().AddCallBack(func(data []byte) {
	// 				sum := uint32(0)
	// 				for i := 0; i < len(data); i += 2 {
	// 					val := uint32(data[i]) | uint32(data[i+1])<<8
	// 					sum += val
	// 				}
	// 				if len(speakerData) > MaxSpeakerDataLength {
	// 					speakerData = speakerData[:0]
	// 				}
	// 				val := sum / uint32((len(data) / 2))
	// 				speakerData = append(speakerData, val)
	// 			})
	// 		}
	// 	}
	// }
	// imgui.End()
}

func drawMainWin() {
	if context.ConnectState != ConnectSuccess {
		return
	}
	drawLocalParticipant()
	drawRemoteParticipants()
}
