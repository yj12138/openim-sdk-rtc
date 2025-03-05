package base

import (
	// lksdk "github.com/livekit/server-sdk-go/v2"
	pb_common "github.com/openimsdk/openim-rtc/proto/go/common"
	pb_event "github.com/openimsdk/openim-rtc/proto/go/event"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/openimsdk/openim-rtc/sdk"
)

type RoomListener struct {
	RoomHandle *pb_common.FfiOwnedHandle
}

// for room
func (l *RoomListener) OnDisconnectedWithReason(reason pb_common.DisconnectReason) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_Disconnected{
			Disconnected: &pb_room.Disconnected{
				Reason: pb_common.DisconnectReason_CLIENT_INITIATED,
			},
		},
	})
}
func (l *RoomListener) OnParticipantConnected(rp *sdk.RemoteParticipant) {
	handle := api.storeObj(rp)
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantConnected{
			ParticipantConnected: &pb_room.ParticipantConnected{
				Info: &pb_participant.OwnedParticipant{
					Handle: handle,
					Info:   &pb_participant.ParticipantInfo{},
				},
			},
		},
	})
}
func (l *RoomListener) OnParticipantDisconnected(rp *sdk.RemoteParticipant) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantDisconnected{
			ParticipantDisconnected: &pb_room.ParticipantDisconnected{},
		},
	})
}
func (l *RoomListener) OnActiveSpeakersChanged(participantIdentities []string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ActiveSpeakersChanged{
			ActiveSpeakersChanged: &pb_room.ActiveSpeakersChanged{
				ParticipantIdentities: participantIdentities,
			},
		},
	})
}
func (l *RoomListener) OnRoomMetadataChanged(metadata string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_RoomMetadataChanged{
			RoomMetadataChanged: &pb_room.RoomMetadataChanged{
				Metadata: metadata,
			},
		},
	})
}
func (l *RoomListener) OnReconnecting() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_Reconnecting{
			Reconnecting: &pb_room.Reconnecting{},
		},
	})
}
func (l *RoomListener) OnReconnected() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_Reconnected{
			Reconnected: &pb_room.Reconnected{},
		},
	})
}

// for local participants
func (l *RoomListener) OnLocalTrackPublished(trackSid string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_LocalTrackPublished{
			LocalTrackPublished: &pb_room.LocalTrackPublished{
				TrackSid: trackSid,
			},
		},
	})
}
func (l *RoomListener) OnLocalTrackUnpublished(publicationSid string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_LocalTrackUnpublished{
			LocalTrackUnpublished: &pb_room.LocalTrackUnpublished{
				PublicationSid: publicationSid,
			},
		},
	})
}

// for all participants
func (l *RoomListener) OnParticipantTrackMuted(participantIdentify string, trackSid string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackMuted{
			TrackMuted: &pb_room.TrackMuted{
				ParticipantIdentity: participantIdentify,
				TrackSid:            trackSid,
			},
		},
	})
}
func (l *RoomListener) OnParticipantTrackUnmuted(participantIdentify string, trackSid string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackUnmuted{
			TrackUnmuted: &pb_room.TrackUnmuted{
				ParticipantIdentity: participantIdentify,
				TrackSid:            trackSid,
			},
		},
	})
}
func (l *RoomListener) OnParticipantMetadataChanged(participantIdentify string, metadata string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantMetadataChanged{
			ParticipantMetadataChanged: &pb_room.ParticipantMetadataChanged{
				ParticipantIdentity: participantIdentify,
				Metadata:            metadata,
			},
		},
	})
}
func (l *RoomListener) OnParticipantAttributesChanged(participantIdentify string, changed map[string]string) {
	changedAttributes := make([]*pb_common.AttributesEntry, 0)
	for k, v := range changed {
		changedAttributes = append(changedAttributes, &pb_common.AttributesEntry{
			Key:   k,
			Value: v,
		})
	}
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantAttributesChanged{
			ParticipantAttributesChanged: &pb_room.ParticipantAttributesChanged{
				ParticipantIdentity: participantIdentify,
				ChangedAttributes:   changedAttributes,
			},
		},
	})
}
func (l *RoomListener) OnIsSpeakingChanged(participantIdentities []string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ActiveSpeakersChanged{
			ActiveSpeakersChanged: &pb_room.ActiveSpeakersChanged{
				ParticipantIdentities: participantIdentities,
			},
		},
	})
}
func (l *RoomListener) OnConnectionQualityChanged(participantIdentify string, quality pb_common.ConnectionQuality) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ConnectionQualityChanged{
			ConnectionQualityChanged: &pb_room.ConnectionQualityChanged{
				ParticipantIdentity: participantIdentify,
				Quality:             quality,
			},
		},
	})
}

// for remote participants
func (l *RoomListener) OnTrackSubscribed(participantIdentify string, track *sdk.RemoteTrack) {
	info := &pb_track.TrackInfo{
		Sid:         track.Publication.SID(),
		Name:        track.Publication.Name(),
		Kind:        sdk.ConvertTrackKind(track.Publication.Kind()),
		StreamState: pb_track.StreamState_STATE_UNKNOWN,
		Muted:       track.Publication.IsMuted(),
		Remote:      true,
	}

	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackSubscribed{
			TrackSubscribed: &pb_room.TrackSubscribed{
				ParticipantIdentity: participantIdentify,
				Track:               info,
			},
		},
	})
}
func (l *RoomListener) OnTrackUnsubscribed(participantIdentify string, trackSid string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackUnsubscribed{
			TrackUnsubscribed: &pb_room.TrackUnsubscribed{
				ParticipantIdentity: participantIdentify,
				TrackSid:            trackSid,
			},
		},
	})
}
func (l *RoomListener) OnTrackSubscriptionFailed(participantIdentify string, trackSid string, err string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackSubscriptionFailed{
			TrackSubscriptionFailed: &pb_room.TrackSubscriptionFailed{
				ParticipantIdentity: participantIdentify,
				TrackSid:            trackSid,
				Error:               err,
			},
		},
	})
}

func (l *RoomListener) OnTrackPublished(participantIdentify string, remoteTrack *sdk.RemoteTrack) {

	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackPublished{
			TrackPublished: &pb_room.TrackPublished{
				ParticipantIdentity: participantIdentify,
				Publication: &pb_track.TrackPublicationInfo{
					Sid:         remoteTrack.Publication.SID(),
					Name:        remoteTrack.Publication.Name(),
					Kind:        sdk.ConvertTrackKind(remoteTrack.Publication.Kind()),
					Source:      pb_track.TrackSource(remoteTrack.Publication.Source()),
					Simulcasted: remoteTrack.Publication.TrackInfo().GetSimulcast(),
					Width:       remoteTrack.Publication.TrackInfo().Height,
					Height:      remoteTrack.Publication.TrackInfo().Height,
					MimeType:    remoteTrack.Publication.MimeType(),
					Muted:       remoteTrack.Publication.IsMuted(),
					Remote:      true,
				},
			},
		},
	})
}

func (l *RoomListener) OnTrackUnpublished(participantIdentify string, publicationSid string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackUnpublished{
			TrackUnpublished: &pb_room.TrackUnpublished{
				ParticipantIdentity: participantIdentify,
				PublicationSid:      publicationSid,
			},
		},
	})
}

func (l *RoomListener) OnDataPacket(participantIdentify string, value interface{}) {
	msg := &pb_room.RoomEvent_DataPacketReceived{
		DataPacketReceived: &pb_room.DataPacketReceived{
			ParticipantIdentity: participantIdentify,
			Value:               nil,
		},
	}
	if user, ok := value.(*pb_room.UserPacket); ok {
		msg.DataPacketReceived.Value = &pb_room.DataPacketReceived_User{
			User: user,
		}
	}
	if sipDtmf, ok := value.(*pb_room.SipDTMF); ok {
		msg.DataPacketReceived.Value = &pb_room.DataPacketReceived_SipDtmf{
			SipDtmf: sipDtmf,
		}
	}
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message:    msg,
	})
}

func (l *RoomListener) OnTranscriptionReceived(participantIdentify string, trackSid string, segments []*pb_participant.TranscriptionSegment) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TranscriptionReceived{
			TranscriptionReceived: &pb_room.TranscriptionReceived{
				ParticipantIdentity: participantIdentify,
				TrackSid:            trackSid,
				Segments:            segments,
			},
		},
	})
}

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}
