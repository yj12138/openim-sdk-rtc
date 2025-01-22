package io

import (
	"fmt"
	"log"

	"github.com/gen2brain/malgo"
)

type Speaker struct {
	context     *malgo.AllocatedContext
	device      *malgo.Device
	canUse      bool
	using       bool
	sizeInBytes uint32

	audioData chan []byte
}

func (m *Speaker) init() {
	m.canUse = false
	m.using = false
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
	})
	if err != nil {
		log.Panic(err)
		return
	}
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.DeviceType = malgo.Capture
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1
	m.sizeInBytes = uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	captureCallbacks := malgo.DeviceCallbacks{
		Data: m.onSendFrames,
		Stop: m.onStop,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		log.Panic(err)
		return
	}
	m.device = device
	m.context = ctx
	m.canUse = true
}

func (m *Speaker) CanUse() bool {
	return m.canUse
}

func (m *Speaker) Using() bool {
	return m.using
}

func (m *Speaker) Start() error {
	if m.canUse {
		err := m.device.Start()
		if err != nil {
			return err
		}
		m.using = true
	}
	return fmt.Errorf("no init device")
}

func (m *Speaker) Stop() error {
	if m.canUse {
		err := m.device.Stop()
		return err
	}
	return fmt.Errorf("no init device")
}

func (m *Speaker) onSendFrames(outputSample, inputSample []byte, framecount uint32) {
	data := <-m.audioData
	copy(outputSample, data)
}

func (m *Speaker) onStop() {
	m.using = false
}

func (m *Speaker) Dispose() {
	if m.device != nil {
		m.device.Uninit()
		_ = m.context.Uninit()
		m.context.Free()
		m.device = nil
	}
}

func (m *Speaker) WriteData(data []byte) {
	m.audioData <- data
}

func NewSpeaker() *Speaker {
	speaker := &Speaker{
		audioData: make(chan []byte, 10),
	}
	speaker.init()
	return speaker
}
