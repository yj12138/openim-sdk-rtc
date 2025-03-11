package sdk

import (
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
)

func ConvertTrackKind(kind lksdk.TrackKind) pb_track.TrackKind {
	if kind == lksdk.TrackKindVideo {
		return pb_track.TrackKind_KIND_VIDEO
	} else if kind == lksdk.TrackKindAudio {
		return pb_track.TrackKind_KIND_AUDIO
	}
	return pb_track.TrackKind_KIND_UNKNOWN
}

func ConvertTrackSource(source livekit.TrackSource) pb_track.TrackSource {
	return pb_track.TrackSource(source)
}
