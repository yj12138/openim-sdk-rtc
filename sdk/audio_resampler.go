package sdk

import (
	"github.com/openimsdk/openim-rtc/sdk/audio"
)

type AudioResampler struct {
}

func (s *AudioResampler) RemixAndResample(sourceData []byte, sourceSampleRate uint32, sourceNumChannels uint32, targetSampleRate uint32, targetChannels uint32) []byte {
	sampleData := audio.GetAudioDSP().Resample(sourceData, sourceSampleRate, sourceNumChannels, targetSampleRate, targetChannels)
	return sampleData
}

func NewAudioResampler() *AudioResampler {
	return &AudioResampler{}
}
