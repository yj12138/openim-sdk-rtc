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

type RemoteTrack struct {
	Track       *webrtc.TrackRemote
	Participant *lksdk.RemoteParticipant
	Publication *lksdk.RemoteTrackPublication
}

func (rt *RemoteTrack) SetEnable(enable bool) {
	rt.Publication.SetEnabled(enable)
}
func (rt *RemoteTrack) SetSubscribed(publicationId string, subscribe bool) {
	rt.Publication.SetSubscribed(subscribe)
}

func (rt *RemoteTrack) EnableTrackPubliciation(remotetrackPublicationSid string, enable bool) {
	rt.Publication.SetEnabled(enable)
}

func (rt *RemoteTrack) UpdatePublicationDimension(remotetrackPublicationSid string, width uint32, height uint32) {
	rt.Publication.SetVideoDimensions(width, height)
}

func NewRemoteTrack(track *webrtc.TrackRemote, remoteParticipant *lksdk.RemoteParticipant, remoteTrackPublication *lksdk.RemoteTrackPublication) *RemoteTrack {
	return &RemoteTrack{
		Track:       track,
		Participant: remoteParticipant,
		Publication: remoteTrackPublication,
	}
}

type LocalTrack struct {
	Name                  string
	MimeType              string
	VideoWidth            int
	VideoHeight           int
	SampleProvider        *RealSampleProvider
	LiveKitTrack          *lksdk.LocalTrack
	LocalTrackPublication *lksdk.LocalTrackPublication
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

func ConvertTrackKind(kind lksdk.TrackKind) pb_track.TrackKind {
	if kind == lksdk.TrackKindVideo {
		return pb_track.TrackKind_KIND_VIDEO
	} else if kind == lksdk.TrackKindAudio {
		return pb_track.TrackKind_KIND_AUDIO
	}
	return pb_track.TrackKind_KIND_UNKNOWN
}

func NewVideoTrack() *LocalTrack {
	return nil
}
