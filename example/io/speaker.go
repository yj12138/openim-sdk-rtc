package io

import (
	"fmt"
	"log"

	"github.com/gen2brain/malgo"
)

type SpeakerCallBackFunc func(data []byte)

type SpeakerCallBackList struct {
	callbacks []SpeakerCallBackFunc
	exists    map[string]bool
}

func (cl *SpeakerCallBackList) Add(callback SpeakerCallBackFunc) {
	callbackID := fmt.Sprintf("%p", callback)
	if !cl.exists[callbackID] {
		cl.callbacks = append(cl.callbacks, callback)
		cl.exists[callbackID] = true
	} else {
		log.Println("回调函数已经存在")
	}
}

func (cl *SpeakerCallBackList) Execute(data []byte) {
	for _, callback := range cl.callbacks {
		callback(data)
	}
}
func newSpeakerCallbackList() *SpeakerCallBackList {
	return &SpeakerCallBackList{
		callbacks: []SpeakerCallBackFunc{},
		exists:    make(map[string]bool),
	}
}

type Speaker struct {
	context *malgo.AllocatedContext
	device  *malgo.Device
	canUse  bool
	using   bool

	audioData chan []byte
	callbacks *SpeakerCallBackList
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
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1
	playbackCallbacks := malgo.DeviceCallbacks{
		Data: m.onSendFrames,
		Stop: m.onStop,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, playbackCallbacks)
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
		return nil
	} else {
		return fmt.Errorf("device not can use")
	}
}

func (m *Speaker) Stop() error {
	if m.canUse {
		err := m.device.Stop()
		return err
	} else {
		return fmt.Errorf("device not can use")
	}
}

func (m *Speaker) onSendFrames(outputSample, inputSample []byte, framecount uint32) {
	data := <-m.audioData
	copy(outputSample, data)
}

func (m *Speaker) onStop() {
	log.Println("Speaker OnStop")
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
	m.callbacks.Execute(data)
	m.audioData <- data
}

func (m *Speaker) AddCallBack(cb func(data []byte)) {
	m.callbacks.Add(cb)
}

func NewSpeaker() *Speaker {
	speaker := &Speaker{
		audioData: make(chan []byte, 10),
		callbacks: newSpeakerCallbackList(),
	}
	speaker.init()
	return speaker
}
