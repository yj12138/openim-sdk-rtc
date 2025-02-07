package sdk

import (
	"context"
	"github.com/pion/webrtc/v4/pkg/media"
	"log"
	"time"
)

type SampleData struct {
	Duration time.Duration
	Data     []byte
}

type RealSampleProvider struct {
	mime            string
	onWriteComplete func()
	dataCache       chan SampleData
}

func (s *RealSampleProvider) OnBind() error {
	log.Println("RealSampleProvider OnBind")
	return nil
}

func (s *RealSampleProvider) OnUnbind() error {
	log.Println("RealSampleProvider OnUnbind")
	return nil
}

func (s *RealSampleProvider) Close() error {
	log.Println("RealSampleProvider Close")
	close(s.dataCache)
	return nil
}

func (s *RealSampleProvider) CurrentAudioLevel() uint8 {
	// default audio level 15
	return 15
}

func (s *RealSampleProvider) WriteData(data []byte, duration time.Duration) {
	s.dataCache <- SampleData{
		Duration: duration,
		Data:     data,
	}
}

func (s *RealSampleProvider) NextSample(c context.Context) (media.Sample, error) {
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

func NewRealSampleProvider(mimeType string) *RealSampleProvider {
	provider := &RealSampleProvider{
		mime:      mimeType,
		dataCache: make(chan SampleData, 10),
	}
	return provider
}
