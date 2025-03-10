package sdk

import (
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
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

func ConvertDataPacket(participantIdentify string, packet *livekit.DataPacket) *pb_room.DataPacketReceived {
	dataPacketReceived := &pb_room.DataPacketReceived{
		ParticipantIdentity: participantIdentify,
		Value:               nil,
	}
	if user, ok := packet.Value.(*livekit.DataPacket_User); ok {
		dataPacketReceived.Value = &pb_room.DataPacketReceived_User{
			User: &pb_room.UserPacket{
				Topic: *user.User.Topic,
				// Data:
			},
		}
	}
	if sipDtmf, ok := packet.Value.(*livekit.DataPacket_SipDtmf); ok {
		dataPacketReceived.Value = &pb_room.DataPacketReceived_SipDtmf{
			SipDtmf: &pb_room.SipDTMF{
				Code:  sipDtmf.SipDtmf.Code,
				Digit: sipDtmf.SipDtmf.Digit,
			},
		}
	}
	return dataPacketReceived
}
