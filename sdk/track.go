package sdk

import (
	"log"
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"
	// pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio"
	// pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	// pb_video "github.com/openimsdk/openim-rtc/proto/go/video"
	"github.com/pion/webrtc/v4"
)

const (
	MimeTypeOpus = "audio/opus"
)

type Track struct {
	name                  string
	mimeType              string
	videoWidth            int
	videoHeight           int
	sampleProvider        *RealSampleProvider
	liveKitTrack          *lksdk.LocalTrack
	localTrackPublication *lksdk.LocalTrackPublication
}

func (t *Track) GetSid() string {
	return t.liveKitTrack.StreamID()
}

func (t *Track) GetName() string {
	return t.name
}

func (t *Track) Kind() pb_track.TrackKind {
	if t.localTrackPublication.Kind() == "video" {
		return pb_track.TrackKind_KIND_VIDEO
	} else if t.localTrackPublication.Kind() == "audio" {
		return pb_track.TrackKind_KIND_AUDIO
	}
	return pb_track.TrackKind_KIND_UNKNOWN
}

func (t *Track) StreamState() pb_track.StreamState {
	return pb_track.StreamState_STATE_ACTIVE
}

func (t *Track) SetMuted(muted bool) {
	t.localTrackPublication.SetMuted(muted)
}

func (t *Track) IsMuted() bool {
	return t.localTrackPublication.IsMuted()
}

func (t *Track) IsRemote() bool {
	return false
}

func (t *Track) Close() {

}

func (t *Track) WriteData(data []byte, duration time.Duration) {
	t.sampleProvider.WriteData(data, duration)
}

func NewAudioTrack(name string) *Track {
	mimeType := MimeTypeOpus
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
