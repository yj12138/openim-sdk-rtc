package sdk

import (
	"log"
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
)

const (
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

func (t *Track) Close() {

}

func (t *Track) WriteData(data []byte, duration time.Duration) {
	t.sampleProvider.WriteData(data, duration)
}

func (t *Track) IsOpened() bool {
	return t.opened
}

func NewAudioTrack(name string, mimeType string) *Track {
	provider := NewRealSampleProvider(mimeType)
	liveKitTrack, err := lksdk.NewLocalTrack(webrtc.RTPCodecCapability{
		MimeType: provider.mime,
	})
	if err != nil {
		log.Panic(err.Error())
	}
	liveKitTrack.OnBind(func() {
		if err := liveKitTrack.StartWrite(provider, provider.onWriteComplete); err != nil {
			log.Panic(err.Error())
		}
	})
	track := &Track{
		name:           name,
		mimeType:       mimeType,
		opened:         false,
		liveKitTrack:   liveKitTrack,
		videoWidth:     0,
		videoHeight:    0,
		sampleProvider: provider,
	}
	return track
}

func NewVideoTrack() *Track {
	return nil
}
