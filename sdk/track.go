package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/pion/webrtc/v4"
	"log"
)

const (
	MimeTypeOpus = "audio/opus"
)

type LocalTrack struct {
	*lksdk.LocalTrack
	// MimeType    string
	VideoWidth  int
	VideoHeight int
}

func NewAudioTrack(name string, source *AudioSource) *LocalTrack {
	// mimeType := MimeTypeOpus
	liveKitTrack, err := lksdk.NewLocalTrack(webrtc.RTPCodecCapability{
		MimeType: source.mime,
	})
	if err != nil {
		log.Panic(err.Error())
	}
	liveKitTrack.OnBind(func() {
		if err := liveKitTrack.StartWrite(source, source.onWriteComplete); err != nil {
			log.Panic(err.Error())
		}
	})
	track := &LocalTrack{
		LocalTrack: liveKitTrack,
		// MimeType:    mimeType,
		VideoWidth:  0,
		VideoHeight: 0,
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
