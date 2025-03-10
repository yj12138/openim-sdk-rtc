package sdk

import (
// "github.com/faiface/beep"
)

type AudioResampler struct {
}

// RemixAndResample 重新采样音频数据
func (resampler *AudioResampler) RemixAndResample(inputData []byte, sourceSampleRate uint32, targetSampleRate uint32) []byte {
	return inputData
}

func NewAudioResampler() *AudioResampler {
	return &AudioResampler{}
}
