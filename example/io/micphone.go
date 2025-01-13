package io

import (
	"fmt"

	"github.com/gen2brain/malgo"
)

type MicPhone struct {
	context             *malgo.AllocatedContext
	device              *malgo.Device
	opened              bool
	capturedSampleCount uint32
	capturedSamples     []byte
}

func (m *MicPhone) init() {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		panic(err)
	}
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1
	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		sampleCount := framecount * deviceConfig.Capture.Channels * sizeInBytes
		newCapturedSampleCount := m.capturedSampleCount + sampleCount
		m.capturedSamples = append(m.capturedSamples, pSample...)
		m.capturedSampleCount = newCapturedSampleCount
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

func (m *MicPhone) Start() error {
	err := m.device.Start()
	if err != nil {
		return err
	}
	return nil
}

func (m *MicPhone) Dispose() {
	m.device.Uninit()
	_ = m.context.Uninit()
	m.context.Free()
}

func NewMicPhone() *MicPhone {
	micPhone := &MicPhone{
		opened:              false,
		capturedSampleCount: 0,
		capturedSamples:     make([]byte, 0),
	}
	micPhone.init()
	return micPhone
}
