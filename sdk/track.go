package sdk

import (
	"log"
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
)

const (
	MimeTypePCMU = "audio/PCMU"
	MimeTypePCMA = "audio/PCMA"
	MimeTypeOpus = "audio/opus"
)

type Track struct {
	name           string
	mimeType       string
	videoWidth     int
	videoHeight    int
	sampleProvider *RealSampleProvider
	liveKitTrack   *lksdk.LocalTrack
	opened         bool
}

func (t *Track) Open() {
	liveKitTrack, err := lksdk.NewLocalTrack(webrtc.RTPCodecCapability{
		MimeType: t.sampleProvider.mime,
	})
	if err != nil {
		log.Panic(err.Error())
	}
	liveKitTrack.OnBind(func() {
		t.opened = true
		// write audio stream or video stream
		if err := liveKitTrack.StartWrite(t.sampleProvider, t.sampleProvider.onWriteComplete); err != nil {
			log.Panic(err.Error())
		}
	})
	t.liveKitTrack = liveKitTrack
}

func (t *Track) Close() {

}

func (t *Track) WriteData(data []byte, duration time.Duration) {
	t.sampleProvider.WriteData(data)
}

func (t *Track) IsOpened() bool {
	return t.opened
}

func NewAudioTrack(name string, mimeType string) *Track {
	track := &Track{
		name:        name,
		mimeType:    mimeType,
		videoWidth:  0,
		videoHeight: 0,
		opened:      false,
	}
	provider := NewRealSampleProvider(mimeType)
	track.sampleProvider = provider

	return track
}

func NewVideoTrack() *Track {
	return nil
}
