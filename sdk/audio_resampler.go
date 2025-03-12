package sdk

import (
	"log"
)

type AudioResampler struct {
	sampleRate     uint32
	targetChannels uint32
}

func (s *AudioResampler) RemixAndResample(sourceData []byte, sourceSampleRate uint32, sourceNumChannels uint32, sourceSamplesPerChannel uint32, targetSampleRate uint32, targetChannels uint32) []byte {
	log.Println("RemixAndResample", "Source Length", len(sourceData), sourceSampleRate, sourceNumChannels, sourceSamplesPerChannel, targetSampleRate, targetChannels)
	if sourceSampleRate == targetSampleRate && sourceNumChannels == targetChannels {
		return sourceData
	}
	var destData []byte = sourceData
	if sourceNumChannels != targetChannels {
		if sourceNumChannels == 1 && targetChannels == 2 {
			destData = make([]byte, len(sourceData)*2)
			for i := 0; i < len(sourceData); i++ {
				destData[i*2] = sourceData[i]
				destData[i*2+1] = sourceData[i]
			}
		}
		if sourceNumChannels == 2 && targetChannels == 1 {
			destData = make([]byte, len(sourceData)/2)
			for i := 0; i < len(sourceData)/2; i++ {
				destData[i] = sourceData[i*2]
			}
		}
	}
	if sourceSampleRate != targetSampleRate {
		// return Resample(destData, int(sourceSampleRate), int(targetSampleRate))
	} else {
	}
	return destData
}

func NewAudioResampler() *AudioResampler {
	return &AudioResampler{}
}
