package io

import (
	"fmt"

	"github.com/gen2brain/malgo"
)

type Speaker struct {
	context *malgo.AllocatedContext
	device  *malgo.Device
}

func (m *Speaker) init() {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		panic(err)
	}

	defer func() {
	}()

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1

	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		sampleCount := framecount * deviceConfig.Capture.Channels * sizeInBytes
		newCapturedSampleCount := capturedSampleCount + sampleCount
		pCapturedSamples = append(pCapturedSamples, pSample...)
		capturedSampleCount = newCapturedSampleCount
	}
	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		panic(err)
	}
	m.device = device
}

func (m *Speaker) Start() error {
	err := m.device.Start()
	if err != nil {
		return err
	}
	return nil
}

func (m *Speaker) Dispose() {
	m.device.Uninit()
	_ = m.context.Uninit()
	m.context.Free()
}

func NewSpeaker() *Speaker {
	speaker := &Speaker{}
	speaker.init()
	return speaker
}
