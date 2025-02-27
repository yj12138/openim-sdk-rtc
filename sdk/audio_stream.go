package sdk

import (
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio"
)

type AudioStream struct {
	StreamType  pb_audio.AudioStreamType
	SampleRate  uint32
	NumChannels uint32
}

func NewAudioStreamByTrack(track *Track, streamType pb_audio.AudioStreamType, sampleRate uint32, numChannels uint32) *AudioStream {
	return &AudioStream{
		StreamType:  streamType,
		SampleRate:  sampleRate,
		NumChannels: numChannels,
	}
}
