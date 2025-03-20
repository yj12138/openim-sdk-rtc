package audio

import (
	"bytes"
	"encoding/binary"
	"log"
)

var audioDSPInstance AudioDSP = nil

type AudioDSP interface {
	Init() error
	EchoCancellation(source []byte, sampleRate uint32, echo []byte) []byte
	Resample(source []byte, sourceSampleRate uint32, sourceNumChannels uint32, targetSampleRate uint32, targetChannels uint32) []byte
	Destory()
}

type DefaultAudioDSP struct{}

func (d *DefaultAudioDSP) Init() error {
	return nil
}
func (d *DefaultAudioDSP) EchoCancellation(source []byte, sampleRate uint32, echo []byte) []byte {
	return source
}
func (d *DefaultAudioDSP) Resample(source []byte, sourceSampleRate uint32, sourceNumChannels uint32, targetSampleRate uint32, targetChannels uint32) []byte {
	return source
}
func (d *DefaultAudioDSP) Destory() {

}

func bytesToInt16(data []byte) ([]int16, error) {
	buf := bytes.NewReader(data)
	int16Data := make([]int16, len(data)/2)
	err := binary.Read(buf, binary.LittleEndian, &int16Data)
	return int16Data, err
}

func int16ToBytes(data []int16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	return buf.Bytes()
}

func GetAudioDSP() AudioDSP {
	if audioDSPInstance == nil {
		speexDSP := &SpeexDSP{}
		err := speexDSP.Init()
		if err != nil {
			log.Println("SpeexDSP Init Error:", err.Error())
			audioDSPInstance = &DefaultAudioDSP{}
		} else {
			audioDSPInstance = speexDSP
		}
	}
	return audioDSPInstance
}
