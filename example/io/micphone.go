package io

import (
	"fmt"
	"log"

	"github.com/gen2brain/malgo"
)

type MicPhoneCallbackFunc func(data []byte, frameCount uint32)

type MicPhoneCallbackList struct {
	callbacks []MicPhoneCallbackFunc
	exists    map[string]bool
}

func newMicPhoneCallbackList() *MicPhoneCallbackList {
	return &MicPhoneCallbackList{
		callbacks: []MicPhoneCallbackFunc{},
		exists:    make(map[string]bool),
	}
}
func (cl *MicPhoneCallbackList) Add(callback MicPhoneCallbackFunc) {
	callbackID := fmt.Sprintf("%p", callback)
	if !cl.exists[callbackID] {
		cl.callbacks = append(cl.callbacks, callback)
		cl.exists[callbackID] = true
	} else {
		log.Println("回调函数已经存在")
	}
}

func (cl *MicPhoneCallbackList) Execute(data []byte, frameCount uint32) {
	for _, callback := range cl.callbacks {
		callback(data, frameCount)
	}
}

type MicPhone struct {
	context     *malgo.AllocatedContext
	device      *malgo.Device
	canUse      bool
	using       bool
	sizeInBytes uint32

	callbacks *MicPhoneCallbackList

	SampleRate uint32
	Channels   uint32
}

func (m *MicPhone) init() {
	m.canUse = false
	m.using = false
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
	})
	if err != nil {
		log.Panic(err)
		return
	}
	captureDevices, err := ctx.Devices(malgo.Capture)
	if err != nil {
		log.Fatalf("无法获取捕获设备")
		return
	}
	log.Println("捕获设备:")
	for _, device := range captureDevices {
		log.Printf("名称: %s FormatCount :%d 默认 %b", device.Name(), device.FormatCount, device.IsDefault)
		for _, format := range device.Formats {
			log.Println(format.Format, format.Channels, format.SampleRate, format.Flags)
		}
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = m.Channels
	deviceConfig.SampleRate = m.SampleRate
	deviceConfig.Alsa.NoMMap = 1
	m.sizeInBytes = uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	captureCallbacks := malgo.DeviceCallbacks{
		Data: m.OnRecvFrames,
		Stop: m.OnStop,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		log.Println(err)
		return
	}
	m.canUse = true
	m.device = device
	m.context = ctx
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
		return nil
	} else {
		return fmt.Errorf("device not can use")
	}
}

func (m *MicPhone) Stop() error {
	if m.canUse {
		err := m.device.Stop()
		return err
	} else {
		return fmt.Errorf("device not can use")
	}
}

func (m *MicPhone) OnRecvFrames(outputSample, inputSample []byte, framecount uint32) {
	m.callbacks.Execute(inputSample, framecount)
}

func (m *MicPhone) OnStop() {
	log.Println("MicPhone OnStop")
	m.using = false
}

func (m *MicPhone) Dispose() {
	if m.device != nil {
		m.device.Uninit()
		_ = m.context.Uninit()
		m.context.Free()
		m.device = nil
	}
}

func (m *MicPhone) AddCallBack(cb func(data []byte, frameCount uint32)) {
	m.callbacks.Add(cb)
}

func NewMicPhone(sameleRate uint32, chnnels uint32) *MicPhone {
	micPhone := &MicPhone{
		SampleRate: sameleRate,
		Channels:   chnnels,
		callbacks:  newMicPhoneCallbackList(),
	}
	micPhone.init()
	return micPhone
}
