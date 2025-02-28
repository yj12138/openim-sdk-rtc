package sdk

import (
	"log"
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/pion/webrtc/v4"
)

const (
	MimeTypeOpus = "audio/opus"
)

type LocalTrack struct {
	Name                  string
	MimeType              string
	VideoWidth            int
	VideoHeight           int
	SampleProvider        *RealSampleProvider
	LiveKitTrack          *lksdk.LocalTrack
	LocalTrackPublication *lksdk.LocalTrackPublication
}

func (t *LocalTrack) Kind() pb_track.TrackKind {
	if t.LocalTrackPublication.Kind() == "video" {
		return pb_track.TrackKind_KIND_VIDEO
	} else if t.LocalTrackPublication.Kind() == "audio" {
		return pb_track.TrackKind_KIND_AUDIO
	}
	return pb_track.TrackKind_KIND_UNKNOWN
}

func (t *LocalTrack) StreamState() pb_track.StreamState {
	return pb_track.StreamState_STATE_ACTIVE
}

func (t *LocalTrack) WriteData(data []byte, duration time.Duration) {
	t.SampleProvider.WriteData(data, duration)
}

func NewAudioTrack(name string) *LocalTrack {
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
	track := &LocalTrack{
		Name:           name,
		MimeType:       mimeType,
		LiveKitTrack:   liveKitTrack,
		VideoWidth:     0,
		VideoHeight:    0,
		SampleProvider: provider,
	}
	return track
}

func NewVideoTrack() *LocalTrack {
	return nil
}

type RemoteTrack struct {
	LiveKitTrack      *webrtc.TrackRemote
	Publication       *lksdk.RemoteTrackPublication
	RemoteParticipant *lksdk.RemoteParticipant
}

func (t *RemoteTrack) Kind() pb_track.TrackKind {
	if t.Publication.Kind() == "video" {
		return pb_track.TrackKind_KIND_VIDEO
	} else if t.Publication.Kind() == "audio" {
		return pb_track.TrackKind_KIND_AUDIO
	}
	return pb_track.TrackKind_KIND_UNKNOWN
}

func NewRemoteTrack(trackRemote *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, remoteParticipant *lksdk.RemoteParticipant) *RemoteTrack {
	return &RemoteTrack{
		LiveKitTrack:      trackRemote,
		Publication:       publication,
		RemoteParticipant: remoteParticipant,
	}
}
