package sdk

import (
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
)

type AudioSource struct {
	SourceType  pb_audio.AudioSourceType
	SampleRate  uint32
	NumChannels uint32
	QueueSizeMs uint32
}

func (source *AudioSource) CaptureFrame(data []byte) {

}

func (source *AudioSource) ClearBuffer() {

}

func NewAudioSource(sourceType pb_audio.AudioSourceType, sampleRate uint32, numChannels uint32, queueSizeMs uint32) *AudioSource {
	return &AudioSource{
		SourceType:  sourceType,
		SampleRate:  sampleRate,
		NumChannels: numChannels,
		QueueSizeMs: queueSizeMs,
	}
}
