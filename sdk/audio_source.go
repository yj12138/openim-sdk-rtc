package sdk

import (
	"context"
	"log"
	"time"

	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	"github.com/openimsdk/openim-rtc/sdk/audio"
	"github.com/pion/webrtc/v4/pkg/media"
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
	return 15
}

func (s *AudioSource) NextSample(c context.Context) (media.Sample, error) {
	sample := media.Sample{}
	sampleData := <-s.dataCache
	sample.Data = sampleData.Data
	sample.Duration = sampleData.Duration
	//log.Println("NextSample:", len(sample.Data), sample.Duration)
	return sample, nil
}

func (s *AudioSource) calcDuration(sampleCount int) time.Duration {
	return time.Duration((sampleCount * 1e9) / int(s.SampleRate))
}

func (s *AudioSource) CaptureFrame(data []byte, numChannels, sampleRate, samplesPerChannel uint32, echo []byte) []byte {
	frameData := audio.GetAudioDSP().AAAProcess(data, sampleRate, nil)
	s.dataCache <- AudioSampleData{
		Duration: s.calcDuration(int(samplesPerChannel) * int(numChannels)),
		Data:     frameData,
	}
	return frameData
}

func (s *AudioSource) ClearBuffer() {

}

func NewAudioSource(sourceType pb_audio.AudioSourceType, sampleRate uint32, numChannels uint32) *AudioSource {
	return &AudioSource{
		SourceType:  sourceType,
		SampleRate:  sampleRate,
		NumChannels: numChannels,
		mime:        "audio/opus",
		dataCache:   make(chan AudioSampleData, 100),
	}
}
