package base

import (
	pb_event "github.com/openimsdk/openim-rtc/proto/go/event"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type RoomListener struct {
	RoomHandle uint64
}

// for room
func (l *RoomListener) OnDisconnectedWithReason(reason pb_participant.DisconnectReason) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_Disconnected{
			Disconnected: &pb_room.Disconnected{
				Reason: pb_participant.DisconnectReason_CLIENT_INITIATED,
			},
		},
	})
}
func (l *RoomListener) OnParticipantConnected() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantConnected{
			ParticipantConnected: &pb_room.ParticipantConnected{},
		},
	})
}
func (l *RoomListener) OnParticipantDisconnected() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantDisconnected{
			ParticipantDisconnected: &pb_room.ParticipantDisconnected{},
		},
	})
}
func (l *RoomListener) OnActiveSpeakersChanged() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ActiveSpeakersChanged{
			ActiveSpeakersChanged: &pb_room.ActiveSpeakersChanged{},
		},
	})
}
func (l *RoomListener) OnRoomMetadataChanged(metadata string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_RoomMetadataChanged{
			RoomMetadataChanged: &pb_room.RoomMetadataChanged{},
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
func (l *RoomListener) OnLocalTrackPublished() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_LocalTrackPublished{
			LocalTrackPublished: &pb_room.LocalTrackPublished{},
		},
	})
}
func (l *RoomListener) OnLocalTrackUnpublished() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_LocalTrackUnpublished{
			LocalTrackUnpublished: &pb_room.LocalTrackUnpublished{},
		},
	})
}

// for all participants
func (l *RoomListener) OnTrackMuted() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackMuted{
			TrackMuted: &pb_room.TrackMuted{},
		},
	})
}
func (l *RoomListener) OnTrackUnmuted() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackUnmuted{
			TrackUnmuted: &pb_room.TrackUnmuted{},
		},
	})
}
func (l *RoomListener) OnMetadataChanged() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_RoomMetadataChanged{
			RoomMetadataChanged: &pb_room.RoomMetadataChanged{},
		},
	})
}
func (l *RoomListener) OnAttributesChanged() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ParticipantAttributesChanged{
			ParticipantAttributesChanged: &pb_room.ParticipantAttributesChanged{},
		},
	})
}
func (l *RoomListener) OnIsSpeakingChanged() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ActiveSpeakersChanged{
			ActiveSpeakersChanged: &pb_room.ActiveSpeakersChanged{},
		},
	})
}
func (l *RoomListener) OnConnectionQualityChanged() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_ConnectionQualityChanged{
			ConnectionQualityChanged: &pb_room.ConnectionQualityChanged{},
		},
	})
}

// for remote participants
func (l *RoomListener) OnTrackSubscribed() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackSubscribed{
			TrackSubscribed: &pb_room.TrackSubscribed{},
		},
	})
}
func (l *RoomListener) OnTrackUnsubscribed() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackUnsubscribed{
			TrackUnsubscribed: &pb_room.TrackUnsubscribed{},
		},
	})
}
func (l *RoomListener) OnTrackSubscriptionFailed() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackSubscriptionFailed{
			TrackSubscriptionFailed: &pb_room.TrackSubscriptionFailed{},
		},
	})
}

func (l *RoomListener) OnTrackPublished() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackPublished{
			TrackPublished: &pb_room.TrackPublished{},
		},
	})
}
func (l *RoomListener) OnTrackUnpublished() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TrackUnpublished{
			TrackUnpublished: &pb_room.TrackUnpublished{},
		},
	})
}

func (l *RoomListener) OnDataPacket() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_DataPacketReceived{
			DataPacketReceived: &pb_room.DataPacketReceived{},
		},
	})
}

func (l *RoomListener) OnTranscriptionReceived() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_TranscriptionReceived{
			TranscriptionReceived: &pb_room.TranscriptionReceived{},
		},
	})
}

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}
