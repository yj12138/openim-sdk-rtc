package sdk

import (
	"context"
	pb_video "github.com/openimsdk/openim-rtc/proto/go/video_frame"
	"github.com/pion/webrtc/v4/pkg/media"
	"log"
	"time"
)

type VideoSampleData struct {
	Duration time.Duration
	Data     []byte
}

type VideoSource struct {
	SourceType      pb_video.VideoSourceType
	Width           uint32
	Height          uint32
	mime            string
	onWriteComplete func()
	dataCache       chan VideoSampleData
}

func (s *VideoSource) OnBind() error {
	log.Println("RealSampleProvider OnBind")
	return nil
}

func (s *VideoSource) OnUnbind() error {
	log.Println("RealSampleProvider OnUnbind")
	return nil
}

func (s *VideoSource) Close() error {
	log.Println("RealSampleProvider Close")
	close(s.dataCache)
	return nil
}

func (s *VideoSource) CurrentAudioLevel() uint8 {
	// default audio level 15
	return 15
}

func (s *VideoSource) NextSample(c context.Context) (media.Sample, error) {
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

func (s *VideoSource) CaptureFrame(timeStampUs int64, rotation pb_video.VideoRotation, data []byte) error {
	s.dataCache <- VideoSampleData{
		Data: data,
	}
	return nil
}

func (s *VideoSource) ClearBuffer() {

}

func NewVideoSource(sourceType pb_video.VideoSourceType, width uint32, height uint32) *VideoSource {
	return &VideoSource{
		SourceType: sourceType,
		Width:      width,
		Height:     height,
		// TODO
		mime:      "",
		dataCache: make(chan VideoSampleData, 10),
	}
}
