package sdk

import (
	"context"
	"log"
	"time"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

type RealSampleProvider struct {
	Mime            string
	FrameDuration   time.Duration
	OnWriteComplete func()
}

func (p *RealSampleProvider) OnBind() error {
	log.Println("RealSampleProvider OnBind")
	return nil
}

func (p *RealSampleProvider) OnUnbind() error {
	log.Println("RealSampleProvider OnUnbind")
	return nil
}

func (p *RealSampleProvider) Close() error {
	log.Println("RealSampleProvider Close")
	return nil
}

func (p RealSampleProvider) CurrentAudioLevel() uint8 {
	return 0
}

func (p *RealSampleProvider) NextSample(c context.Context) (media.Sample, error) {
	// log.Println("RealSampleProvider NextSample")
	sample := media.Sample{}
	switch p.Mime {
	case webrtc.MimeTypePCMA:
		// sample.Data = pcmaData
		// sample.Duration = duration
	}
	return sample, nil
}
