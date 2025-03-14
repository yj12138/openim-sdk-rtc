package io

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/gen2brain/malgo"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

type MicPhone struct {
	context     *malgo.AllocatedContext
	device      *malgo.Device
	canUse      bool
	using       bool
	sizeInBytes uint32

	SampleRate   uint32
	Channels     uint32
	rawAudioData []byte

	CallBack func([]byte, uint32)
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
	m.rawAudioData = append(m.rawAudioData, inputSample...)
	if m.CallBack != nil {
		m.CallBack(inputSample, framecount)
	}
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

func (m *MicPhone) SaveWavFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteToPCM := func(data []byte) []int {
		pcmData := make([]int, len(data)/2)
		for i := 0; i < len(data); i += 2 {
			pcmData[i/2] = int(binary.LittleEndian.Uint16(data[i : i+2]))
		}
		return pcmData
	}

	enc := wav.NewEncoder(file, int(m.SampleRate), 16, int(m.Channels), 1)

	buf := &audio.IntBuffer{
		Format: &audio.Format{
			SampleRate:  int(m.SampleRate),
			NumChannels: int(m.Channels),
		},
		Data:           byteToPCM(m.rawAudioData),
		SourceBitDepth: 16,
	}

	if err := enc.Write(buf); err != nil {
		return err
	}

	if err := enc.Close(); err != nil {
		return err
	}

	fmt.Println("WAV 文件已保存:", filename)
	return nil
}

func NewMicPhone(sameleRate uint32, chnnels uint32) *MicPhone {
	micPhone := &MicPhone{
		SampleRate:   sameleRate,
		Channels:     chnnels,
		rawAudioData: make([]byte, 0),
	}
	micPhone.init()
	return micPhone
}
