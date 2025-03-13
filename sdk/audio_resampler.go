package sdk

import (
	"encoding/binary"
	"math"
)

type AudioResampler struct {
}

// 将 PCM 字节数据转换为浮点型数组
func bytesToFloat64s(pcmBytes []byte) []float64 {
	sampleCount := len(pcmBytes) / 2
	floatSamples := make([]float64, sampleCount)
	for i := 0; i < sampleCount; i++ {
		sample := int16(binary.LittleEndian.Uint16(pcmBytes[i*2 : (i+1)*2]))
		floatSamples[i] = float64(sample) / math.MaxInt16
	}
	return floatSamples
}

// 移动平均滤波器
func movingAverageFilter(samples []float64, windowSize int) []float64 {
	if windowSize <= 1 {
		return samples
	}
	filteredSamples := make([]float64, len(samples))
	var sum float64
	for i := 0; i < len(samples); i++ {
		sum += samples[i]
		if i >= windowSize {
			sum -= samples[i-windowSize]
		}
		if i >= windowSize-1 {
			filteredSamples[i-windowSize/2] = sum / float64(windowSize)
		}
	}
	return filteredSamples
}

// 将浮点型数组转换回 PCM 字节数据
func float64sToBytes(floatSamples []float64) []byte {
	pcmBytes := make([]byte, len(floatSamples)*2)
	for i, sample := range floatSamples {
		intSample := int16(sample * math.MaxInt16)
		binary.LittleEndian.PutUint16(pcmBytes[i*2:], uint16(intSample))
	}
	return pcmBytes
}

func (s *AudioResampler) RemixAndResample(sourceData []byte, sourceSampleRate uint32, sourceNumChannels uint32, sourceSamplesPerChannel uint32, targetSampleRate uint32, targetChannels uint32) []byte {
	// log.Println("RemixAndResample", "Source Length", len(sourceData), sourceSampleRate, sourceNumChannels, sourceSamplesPerChannel, targetSampleRate, targetChannels)
	// 降噪处理
	floatSamples := bytesToFloat64s(sourceData)
	// 应用移动平均滤波器进行降噪
	windowSize := 5 // 滤波窗口大小
	filteredSamples := movingAverageFilter(floatSamples, windowSize)

	// 将处理后的浮点型数组转换回 PCM 字节数据
	filteredPCMData := float64sToBytes(filteredSamples)
	return filteredPCMData
	// 重采样
	// if sourceSampleRate == targetSampleRate && sourceNumChannels == targetChannels {
	// 	return sourceData
	// }
	// var destData []byte = sourceData
	// if sourceNumChannels != targetChannels {
	// 	if sourceNumChannels == 1 && targetChannels == 2 {
	// 		destData = make([]byte, len(sourceData)*2)
	// 		for i := 0; i < len(sourceData); i++ {
	// 			destData[i*2] = sourceData[i]
	// 			destData[i*2+1] = sourceData[i]
	// 		}
	// 	}
	// 	if sourceNumChannels == 2 && targetChannels == 1 {
	// 		destData = make([]byte, len(sourceData)/2)
	// 		for i := 0; i < len(sourceData)/2; i++ {
	// 			destData[i] = sourceData[i*2]
	// 		}
	// 	}
	// }
	// if sourceSampleRate != targetSampleRate {
	// 	// return Resample(destData, int(sourceSampleRate), int(targetSampleRate))
	// } else {
	// }
	// return destData
}

func NewAudioResampler() *AudioResampler {
	return &AudioResampler{}
}
