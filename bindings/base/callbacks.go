package base

import (
	pb_event "github.com/openimsdk/openim-rtc/proto/go/event"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type RoomListener struct {
	RoomHandle uint64
}

func (l *RoomListener) OnDisconnected() {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_Disconnected{
			Disconnected: &pb_room.Disconnected{},
		},
	})
}
func (l *RoomListener) OnDisconnectedWithReason(reason string) {
	dispatchEventResp(pb_event.FuncEventName_RoomEvent, &pb_room.RoomEvent{
		RoomHandle: l.RoomHandle,
		Message: &pb_room.RoomEvent_Disconnected{
			Disconnected: &pb_room.Disconnected{},
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

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}
