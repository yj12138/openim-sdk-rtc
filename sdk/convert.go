package sdk

import (
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
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

func ConvertParticipantKind(kind lksdk.ParticipantKind) pb_participant.ParticipantKind {
	return pb_participant.ParticipantKind(kind)
}

func ConvertTranscriptionSegment(transcriptionSegments []*lksdk.TranscriptionSegment) []*pb_room.TranscriptionSegment {
	segments := make([]*pb_room.TranscriptionSegment, len(transcriptionSegments))
	for i, segment := range transcriptionSegments {
		segments[i] = &pb_room.TranscriptionSegment{
			Id:        segment.ID,
			Text:      segment.Text,
			StartTime: segment.StartTime,
			EndTime:   segment.EndTime,
			Final:     segment.Final,
			Language:  segment.Language,
		}
	}
	return segments
}
