package sdk

import (
	"context"
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	"github.com/pion/webrtc/v4/pkg/media"
	"log"
	"time"
)

type AudioSampleData struct {
	Duration time.Duration
	Data     []byte
}

type AudioSource struct {
	SourceType      pb_audio.AudioSourceType
	SampleRate      uint32
	NumChannels     uint32
	mime            string
	onWriteComplete func()
	dataCache       chan AudioSampleData
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

func (s *AudioSource) NextSample(c context.Context) (media.Sample, error) {
	sample := media.Sample{}
	select {
	case sampleData := <-s.dataCache:
		sample.Data = sampleData.Data
		sample.Duration = sampleData.Duration
		// log.Println("NextSample:", len(sample.Data), sample.Duration)
	default:
		sample.Data = make([]byte, 0)
		sample.Duration = 1 * time.Second
	}
	return sample, nil
}

func (s *AudioSource) calcDuration(sampleCount int) time.Duration {
	// 每两个字节表示一个采样
	return time.Duration((sampleCount / 2) * 1e9 / int(s.SampleRate))
}

func (s *AudioSource) CaptureFrame(data []byte) error {
	s.dataCache <- AudioSampleData{
		Duration: s.calcDuration(len(data)),
		Data:     data,
	}
	return nil
}

func (s *AudioSource) ClearBuffer() {

}

func NewAudioSource(sourceType pb_audio.AudioSourceType, sampleRate uint32, numChannels uint32) *AudioSource {
	return &AudioSource{
		SourceType:  sourceType,
		SampleRate:  sampleRate,
		NumChannels: numChannels,
		// TODO
		mime:      "audio/opus",
		dataCache: make(chan AudioSampleData, 10),
	}
}
