package sdk

import (
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio"
)

type AudioSource struct {
	SourceType  pb_audio.AudioSourceType
	SampleRate  uint32
	NumChannels uint32
	QueueSizeMs uint32
}

func (source *AudioSource) CaptureFrame() {

}

func NewAudioSource() *AudioSource {
	return &AudioSource{}
}
