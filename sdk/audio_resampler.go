package sdk

import (
	// "encoding/binary"
	"github.com/openimsdk/openim-rtc/sdk/audio"
	// "math"
)

type AudioResampler struct {
}

func (s *AudioResampler) RemixAndResample(sourceData []byte, sourceSampleRate uint32, sourceNumChannels uint32, sourceSamplesPerChannel uint32, targetSampleRate uint32, targetChannels uint32) []byte {
	// TODO
	audio.Resample()
	return sourceData
}

func NewAudioResampler() *AudioResampler {
	return &AudioResampler{}
}
