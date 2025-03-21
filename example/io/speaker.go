package io

import (
	"fmt"
	"log"

	"github.com/gen2brain/malgo"
	"github.com/openimsdk/openim-rtc/example/common"
	"github.com/openimsdk/openim-rtc/sdk"
)

type Speaker struct {
	context *malgo.AllocatedContext
	device  *malgo.Device
	canUse  bool
	using   bool

	sampleRate uint32
	channels   uint32

	resample *sdk.AudioResampler
	curFrame []byte

	audioBuffer *common.RingBuffer
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
	playbackDevices, err := ctx.Devices(malgo.Playback)
	if err != nil {
		log.Fatalf("无法获取播放设备: %v", err)
	}

	// 输出每个播放设备的名称和支持的最大通道数
	log.Println("播放设备:")
	for _, device := range playbackDevices {
		log.Printf("名称: %s FormatCount :%d 默认 %b", device.Name(), device.FormatCount, device.IsDefault)
		for _, format := range device.Formats {
			log.Println(format.Format, format.Channels, format.SampleRate, format.Flags)
		}
	}
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = m.channels
	deviceConfig.SampleRate = m.sampleRate
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
	_, err := m.audioBuffer.Read(outputSample)
	if err != nil {
		return
	}
	if m.curFrame == nil || len(m.curFrame) != len(outputSample) {
		m.curFrame = make([]byte, len(outputSample))
	}
	copy(m.curFrame, outputSample)
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

func (m *Speaker) Write(data []byte) {
	_, err := m.audioBuffer.Write(data)
	if err != nil {
		log.Println("Write Buffer Error:", err.Error())
	}
}

func (m *Speaker) GetLastFrame() []byte {
	return m.curFrame
}

func NewSpeaker(sampleRate uint32, channels uint32) *Speaker {
	speaker := &Speaker{
		sampleRate:  sampleRate,
		channels:    channels,
		resample:    sdk.NewAudioResampler(),
		curFrame:    nil,
		audioBuffer: common.NewRingBuffer(int(float32(sampleRate) * float32(channels) * 0.5)),
	}
	speaker.init()
	return speaker
}
