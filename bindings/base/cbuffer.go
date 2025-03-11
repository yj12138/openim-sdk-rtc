package base

import (
	pb_video "github.com/openimsdk/openim-rtc/proto/go/video_frame"
)

type CBuffer struct {
	SourceData []byte
	CDataPtr   uint64
	CDataLen   int
}

func NewCBuffer(sourceData []byte) *CBuffer {
	ptr := goByteSliceToCPointerNoCopyFunc(sourceData)
	return &CBuffer{
		SourceData: sourceData,
		CDataPtr:   ptr,
		CDataLen:   len(sourceData),
	}
}

type AudioFrameBuffer struct {
	DataPtr           uint64
	NumChannels       uint32
	SampleRate        uint32
	SamplesPerChannel uint32
}

func NewAudioFrameBuffer(data []byte, numChannels uint32, sampleRate uint32, samplesPerChannel uint32) *AudioFrameBuffer {
	return &AudioFrameBuffer{
		DataPtr:           goByteSliceToCPointerNoCopyFunc(data),
		NumChannels:       numChannels,
		SampleRate:        sampleRate,
		SamplesPerChannel: samplesPerChannel,
	}
}

type VideoFrameBuffer struct {
	Type       pb_video.VideoBufferType
	Width      uint32
	Height     uint32
	DataPtr    uint64
	Stride     uint32
	Components []*pb_video.VideoBufferInfo_ComponentInfo
}
