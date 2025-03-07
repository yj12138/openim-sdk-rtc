package base

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/openimsdk/openim-rtc/sdk"
)

type RoomListener struct {
	RoomHandle uint64
}

// for room
func (l *RoomListener) OnDisconnectedWithReason(reason pb_participant.DisconnectReason) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_Disconnected{
				Disconnected: &pb_room.Disconnected{
					Reason: pb_participant.DisconnectReason_CLIENT_INITIATED,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnParticipantConnected(rp *lksdk.RemoteParticipant) {
	handle := api.storeObj(rp)
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ParticipantConnected{
				ParticipantConnected: &pb_room.ParticipantConnected{
					Info: &pb_participant.OwnedParticipant{
						Handle: &pb_handle.FfiOwnedHandle{Id: handle},
						Info:   &pb_participant.ParticipantInfo{},
					},
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ParticipantDisconnected{
				ParticipantDisconnected: &pb_room.ParticipantDisconnected{},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnActiveSpeakersChanged(participantIdentities []string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ActiveSpeakersChanged{
				ActiveSpeakersChanged: &pb_room.ActiveSpeakersChanged{
					ParticipantIdentities: participantIdentities,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

func (l *RoomListener) OnRoomMetadataChanged(metadata string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_RoomMetadataChanged{
				RoomMetadataChanged: &pb_room.RoomMetadataChanged{
					Metadata: metadata,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnReconnecting() {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_Reconnecting{
				Reconnecting: &pb_room.Reconnecting{},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnReconnected() {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_Reconnected{
				Reconnected: &pb_room.Reconnected{},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

// for local participants
func (l *RoomListener) OnLocalTrackPublished(trackSid string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_LocalTrackPublished{
				LocalTrackPublished: &pb_room.LocalTrackPublished{
					TrackSid: trackSid,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnLocalTrackUnpublished(publicationSid string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_LocalTrackUnpublished{
				LocalTrackUnpublished: &pb_room.LocalTrackUnpublished{
					PublicationSid: publicationSid,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

// for all participants
func (l *RoomListener) OnParticipantTrackMuted(participantIdentify string, trackSid string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackMuted{
				TrackMuted: &pb_room.TrackMuted{
					ParticipantIdentity: participantIdentify,
					TrackSid:            trackSid,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnParticipantTrackUnmuted(participantIdentify string, trackSid string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackUnmuted{
				TrackUnmuted: &pb_room.TrackUnmuted{
					ParticipantIdentity: participantIdentify,
					TrackSid:            trackSid,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnParticipantMetadataChanged(participantIdentify string, metadata string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ParticipantMetadataChanged{
				ParticipantMetadataChanged: &pb_room.ParticipantMetadataChanged{
					ParticipantIdentity: participantIdentify,
					Metadata:            metadata,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnParticipantAttributesChanged(participantIdentify string, changed map[string]string) {

	changedAttributes := make([]*pb_room.AttributesEntry, 0)
	for k, v := range changed {
		changedAttributes = append(changedAttributes, &pb_room.AttributesEntry{
			Key:   k,
			Value: v,
		})
	}
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ParticipantAttributesChanged{
				ParticipantAttributesChanged: &pb_room.ParticipantAttributesChanged{
					ParticipantIdentity: participantIdentify,
					ChangedAttributes:   changedAttributes,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnIsSpeakingChanged(participantIdentities []string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ActiveSpeakersChanged{
				ActiveSpeakersChanged: &pb_room.ActiveSpeakersChanged{
					ParticipantIdentities: participantIdentities,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnConnectionQualityChanged(participantIdentify string, quality pb_room.ConnectionQuality) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_ConnectionQualityChanged{
				ConnectionQualityChanged: &pb_room.ConnectionQualityChanged{
					ParticipantIdentity: participantIdentify,
					Quality:             quality,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

// for remote participants
func (l *RoomListener) OnTrackSubscribed(participantIdentify string, publication *lksdk.RemoteTrackPublication) {
	info := &pb_track.TrackInfo{
		Sid:         publication.SID(),
		Name:        publication.Name(),
		Kind:        sdk.ConvertTrackKind(publication.Kind()),
		StreamState: pb_track.StreamState_STATE_UNKNOWN,
		Muted:       publication.IsMuted(),
		Remote:      true,
	}
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackSubscribed{
				TrackSubscribed: &pb_room.TrackSubscribed{
					ParticipantIdentity: participantIdentify,
					Track: &pb_track.OwnedTrack{
						// TODO
						// Handle: &pb_common.FfiOwnedHandle{Id: api.storeObj(track)},
						Info: info,
					},
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnTrackUnsubscribed(participantIdentify string, trackSid string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackUnsubscribed{
				TrackUnsubscribed: &pb_room.TrackUnsubscribed{
					ParticipantIdentity: participantIdentify,
					TrackSid:            trackSid,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}
func (l *RoomListener) OnTrackSubscriptionFailed(participantIdentify string, trackSid string, err string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackSubscriptionFailed{
				TrackSubscriptionFailed: &pb_room.TrackSubscriptionFailed{
					ParticipantIdentity: participantIdentify,
					TrackSid:            trackSid,
					Error:               err,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

func (l *RoomListener) OnTrackPublished(participantIdentify string, publication *lksdk.RemoteTrackPublication) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackPublished{
				TrackPublished: &pb_room.TrackPublished{
					ParticipantIdentity: participantIdentify,
					Publication: &pb_track.OwnedTrackPublication{
						// TODO
						// Handle: &pb_common.FfiOwnedHandle{Id: api.storeObj(remoteTrack)},
						Info: &pb_track.TrackPublicationInfo{
							Sid:         publication.SID(),
							Name:        publication.Name(),
							Kind:        sdk.ConvertTrackKind(publication.Kind()),
							Source:      pb_track.TrackSource(publication.Source()),
							Simulcasted: publication.TrackInfo().GetSimulcast(),
							Width:       publication.TrackInfo().Height,
							Height:      publication.TrackInfo().Height,
							MimeType:    publication.MimeType(),
							Muted:       publication.IsMuted(),
							Remote:      true,
						},
					},
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

func (l *RoomListener) OnTrackUnpublished(participantIdentify string, publicationSid string) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TrackUnpublished{
				TrackUnpublished: &pb_room.TrackUnpublished{
					ParticipantIdentity: participantIdentify,
					PublicationSid:      publicationSid,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

func (l *RoomListener) OnDataPacket(participantIdentify string, value interface{}) {
	message := &pb_room.RoomEvent_DataPacketReceived{
		DataPacketReceived: &pb_room.DataPacketReceived{
			ParticipantIdentity: participantIdentify,
			Value:               nil,
		},
	}
	if user, ok := value.(*pb_room.UserPacket); ok {
		message.DataPacketReceived.Value = &pb_room.DataPacketReceived_User{
			User: user,
		}
	}
	if sipDtmf, ok := value.(*pb_room.SipDTMF); ok {
		message.DataPacketReceived.Value = &pb_room.DataPacketReceived_SipDtmf{
			SipDtmf: sipDtmf,
		}
	}
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message:    message,
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

func (l *RoomListener) OnTranscriptionReceived(participantIdentify string, trackSid string, segments []*pb_room.TranscriptionSegment) {
	msg := &pb_ffi.FfiEvent_RoomEvent{
		RoomEvent: &pb_room.RoomEvent{
			RoomHandle: l.RoomHandle,
			Message: &pb_room.RoomEvent_TranscriptionReceived{
				TranscriptionReceived: &pb_room.TranscriptionReceived{
					ParticipantIdentity: participantIdentify,
					TrackSid:            trackSid,
					Segments:            segments,
				},
			},
		},
	}
	dispatchEvent(&pb_ffi.FfiEvent{Message: msg})
}

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}
