package ui

import (
	"log"

	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type RoomListener struct {
	RoomHandle uint64
}

// for room
func (l *RoomListener) OnDisconnectedWithReason(reason pb_participant.DisconnectReason) {
	log.Println("OnDisconnectedWithReason", reason)
}
func (l *RoomListener) OnParticipantConnected(rp *lksdk.RemoteParticipant) {
	log.Println("OnParticipantConnected", rp.Identity())
}
func (l *RoomListener) OnParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	log.Println("OnParticipantDisconnected", rp.Identity())
}
func (l *RoomListener) OnActiveSpeakersChanged(participantIdentities []string) {
	log.Println("OnActiveSpeakersChanged", participantIdentities)
}

func (l *RoomListener) OnRoomMetadataChanged(metadata string) {
	log.Println("OnRoomMetadataChanged", metadata)
}
func (l *RoomListener) OnReconnecting() {
	log.Println("OnReconnecting")
}
func (l *RoomListener) OnReconnected() {
	log.Println("OnReconnected")
}

// for local participants
func (l *RoomListener) OnLocalTrackPublished(trackSid string) {
	log.Println("OnLocalTrackPublished", trackSid)
}
func (l *RoomListener) OnLocalTrackUnpublished(publicationSid string) {
	log.Println("OnLocalTrackUnpublished", publicationSid)
}

// for all participants
func (l *RoomListener) OnParticipantTrackMuted(participantIdentify string, trackSid string) {
	log.Println("OnParticipantTrackMuted", participantIdentify, trackSid)
}
func (l *RoomListener) OnParticipantTrackUnmuted(participantIdentify string, trackSid string) {
	log.Println("OnParticipantTrackUnmuted", participantIdentify, trackSid)
}
func (l *RoomListener) OnParticipantMetadataChanged(participantIdentify string, metadata string) {
	log.Println("OnParticipantMetadataChanged", participantIdentify, metadata)
}
func (l *RoomListener) OnParticipantAttributesChanged(participantIdentify string, changed map[string]string) {
	log.Println("OnParticipantAttributesChanged", participantIdentify, changed)
}
func (l *RoomListener) OnIsSpeakingChanged(participantIdentities []string) {
	log.Println("OnIsSpeakingChanged", participantIdentities)
}
func (l *RoomListener) OnConnectionQualityChanged(participantIdentify string, quality pb_room.ConnectionQuality) {
	log.Println("OnConnectionQualityChanged", participantIdentify, quality)
}

// for remote participants
func (l *RoomListener) OnTrackSubscribed(participantIdentify string, publication *lksdk.RemoteTrackPublication) {
	log.Println("OnTrackSubscribed", participantIdentify, publication.Name())
}
func (l *RoomListener) OnTrackUnsubscribed(participantIdentify string, trackSid string) {
	log.Println("OnTrackUnsubscribed", participantIdentify, trackSid)
}
func (l *RoomListener) OnTrackSubscriptionFailed(participantIdentify string, trackSid string, err string) {
	log.Println("OnTrackSubscriptionFailed", participantIdentify, trackSid, err)
}

func (l *RoomListener) OnTrackPublished(participantIdentify string, publication *lksdk.RemoteTrackPublication) {
	log.Println("OnTrackPublished", participantIdentify, publication.Name())
}

func (l *RoomListener) OnTrackUnpublished(participantIdentify string, publicationSid string) {
	log.Println("OnTrackUnpublished", participantIdentify, publicationSid)
}

func (l *RoomListener) OnDataPacket(participantIdentify string, packet lksdk.DataPacket) {
	log.Println("OnDataPacket", participantIdentify, packet)
	// packet := data.ToProto()
	// var value *livekit.DataPacket = nil
	// if user, ok := packet.Value.(*livekit.DataPacket_User); ok {
	// 	value = &pb_room.UserPacket{
	// 		Topic: *user.User.Topic,
	// 		// TODO
	// 		// Data:  user.User.Payload,
	// 	}
	// }
	// if sipDtmf, ok := packet.Value.(*livekit.DataPacket_SipDtmf); ok {
	// 	value = &pb_room.SipDTMF{
	// 		Digit: sipDtmf.SipDtmf.Digit,
	// 		Code:  sipDtmf.SipDtmf.Code,
	// 	}
	// }
}

func (l *RoomListener) OnTranscriptionReceived(participantIdentify string, trackSid string, segments []*pb_room.TranscriptionSegment) {
	log.Println("OnTranscriptionReceived", participantIdentify, trackSid, segments)
}

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}
