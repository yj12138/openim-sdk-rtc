package sdk

import (
	"context"
	"log"
	"time"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

type RealSampleProvider struct {
	mime            string
	onWriteComplete func()
	frameDuration   time.Duration
	dataCache       chan []byte
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

func (s *RealSampleProvider) WriteData(data []byte) {
	s.dataCache <- data
}

func (s *RealSampleProvider) NextSample(c context.Context) (media.Sample, error) {
	sample := media.Sample{}
	switch s.mime {
	case webrtc.MimeTypePCMA:
		select {
		case data := <-s.dataCache:
			sample.Data = data
			sample.Duration = s.frameDuration
		default:
			sample.Data = make([]byte, 0)
			sample.Duration = s.frameDuration
		}
	}
	return sample, nil
}

func NewRealSampleProvider(mimeType string) *RealSampleProvider {
	provider := &RealSampleProvider{
		mime:      mimeType,
		dataCache: make(chan []byte, 100),
	}
	return provider
}
