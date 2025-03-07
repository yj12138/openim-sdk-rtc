package sdk

import (
	"context"
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	"github.com/pion/webrtc/v4/pkg/media"
	"log"
	"time"
)

type SampleData struct {
	Duration time.Duration
	Data     []byte
}

type AudioSource struct {
	SourceType      pb_audio.AudioSourceType
	SampleRate      uint32
	NumChannels     uint32
	QueueSizeMs     uint32
	mime            string
	onWriteComplete func()
	dataCache       chan SampleData
}

func (s *AudioSource) OnBind() error {
	log.Println("RealSampleProvider OnBind")
	return nil
}

func (s *AudioSource) OnUnbind() error {
	log.Println("RealSampleProvider OnUnbind")
	return nil
}

func (s *AudioSource) Close() error {
	log.Println("RealSampleProvider Close")
	close(s.dataCache)
	return nil
}

func (s *AudioSource) CurrentAudioLevel() uint8 {
	// default audio level 15
	return 15
}

func (s *AudioSource) WriteData(data []byte, duration time.Duration) {
	s.dataCache <- SampleData{
		Duration: duration,
		Data:     data,
	}
}

func (s *AudioSource) NextSample(c context.Context) (media.Sample, error) {
	sample := media.Sample{}
	switch s.mime {
	case MimeTypeOpus:
		select {
		case sampleData := <-s.dataCache:
			sample.Data = sampleData.Data
			sample.Duration = sampleData.Duration
		default:
			sample.Data = make([]byte, 0)
			sample.Duration = 1 * time.Second
		}
	}
	return sample, nil
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
		// TODO
		mime:      MimeTypeOpus,
		dataCache: make(chan SampleData, 10),
	}
}
