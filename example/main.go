package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/AllenDang/cimgui-go/backend"
	// ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	// "github.com/AllenDang/cimgui-go/imgui"
	"github.com/gen2brain/malgo"
)

// var currentBackend backend.Backend[ebitenbackend.EbitenBackendFlags]

func init() {
	log.SetFlags(log.Llongfile)
}

func loop() {
	// ui.DrawMainMenu()
	// ui.DrawAllWindow()
}

func main() {
	// userId := flag.String("userId", "", "User ID ")
	// token := flag.String("token", "", "User Token")
	// flag.Parse()

	// if *userId == "" || *token == "" {
	// 	flag.Usage()
	// 	return
	// }

	// nickName := ui.Login(*userId, *token)

	// currentBackend, _ = backend.CreateBackend(ebitenbackend.NewEbitenBackend())
	// currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	// currentBackend.CreateWindow("OpenIM-RTC-Demo", 1200, 900)
	// currentBackend.SetCloseCallback(func(b backend.Backend[ebitenbackend.EbitenBackendFlags]) {
	// })

	// currentBackend.Run(loop)

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		// fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1

	var playbackSampleCount uint32
	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		sampleCount := framecount * deviceConfig.Capture.Channels * sizeInBytes
		newCapturedSampleCount := capturedSampleCount + sampleCount
		pCapturedSamples = append(pCapturedSamples, pSample...)
		capturedSampleCount = newCapturedSampleCount
	}

	fmt.Println("Recording...")
	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Press Enter to stop recording...")
	fmt.Scanln()

	device.Uninit()

	onSendFrames := func(pSample, nil []byte, framecount uint32) {
		samplesToRead := framecount * deviceConfig.Playback.Channels * sizeInBytes
		if samplesToRead > capturedSampleCount-playbackSampleCount {
			samplesToRead = capturedSampleCount - playbackSampleCount
		}
		copy(pSample, pCapturedSamples[playbackSampleCount:playbackSampleCount+samplesToRead])
		playbackSampleCount += samplesToRead
		if playbackSampleCount == uint32(len(pCapturedSamples)) {
			playbackSampleCount = 0
		}
	}

	fmt.Println("Playing...")
	playbackCallbacks := malgo.DeviceCallbacks{
		Data: onSendFrames,
	}
	device, err = malgo.InitDevice(ctx.Context, deviceConfig, playbackCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Press Enter to quit...")
	fmt.Scanln()

	device.Uninit()

}
