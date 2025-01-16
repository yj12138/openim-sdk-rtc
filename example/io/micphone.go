package io

import (
	"fmt"
	"log"

	"github.com/gen2brain/malgo"
)

type MicPhone struct {
	context     *malgo.AllocatedContext
	device      *malgo.Device
	canUse      bool
	using       bool
	sizeInBytes uint32
}

func (m *MicPhone) init() {
	m.canUse = false
	m.using = false
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
	})
	if err != nil {
		log.Panic(err)
	}
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1
	m.sizeInBytes = uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	captureCallbacks := malgo.DeviceCallbacks{
		Data: m.OnRecvFrames,
		Stop: m.OnStop,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		log.Panic(err)
	}
	m.canUse = true
	m.device = device
}

func (m *MicPhone) CanUse() bool {
	return m.canUse
}

func (m *MicPhone) Using() bool {
	return m.using
}

func (m *MicPhone) Start() error {
	if m.canUse {
		err := m.device.Start()
		if err != nil {
			return err
		}
		m.using = true
	}
	return fmt.Errorf("no init device")
}

func (m *MicPhone) Stop() error {
	if m.canUse {
		err := m.device.Stop()
		return err
	}
	return fmt.Errorf("no init device")
}

func (m *MicPhone) OnRecvFrames(outputSample, inputSample []byte, framecount uint32) {

}
func (m *MicPhone) OnStop() {
	m.using = false
}

func (m *MicPhone) Dispose() {
	m.device.Uninit()
	_ = m.context.Uninit()
	m.context.Free()
}

func NewMicPhone() *MicPhone {
	micPhone := &MicPhone{}
	micPhone.init()
	return micPhone
}
