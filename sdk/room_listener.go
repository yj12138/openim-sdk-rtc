package sdk

import (
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type OnRoomListener interface {
	// room
	OnDisconnectedWithReason(pb_participant.DisconnectReason)
	OnParticipantConnected(*RemoteParticipant)
	OnParticipantDisconnected(*RemoteParticipant)
	OnActiveSpeakersChanged([]string)
	OnRoomMetadataChanged(metadata string)
	OnReconnecting()
	OnReconnected()
	// for local participants
	OnLocalTrackPublished(trackSid string)
	OnLocalTrackUnpublished(publicationSid string)

	// for all participants
	OnParticipantTrackMuted(participantIdentify string, trackSid string)
	OnParticipantTrackUnmuted(participantIdentify string, trackSid string)
	OnParticipantMetadataChanged(participantIdentify string, metadata string)
	OnParticipantAttributesChanged(participantIdentify string, changed map[string]string)
	OnIsSpeakingChanged(participantIdentify []string)
	OnConnectionQualityChanged(participantIdentify string, quality pb_room.ConnectionQuality)

	// for remote participants
	OnTrackSubscribed(participantIdentify string, track *RemoteTrack)
	OnTrackUnsubscribed(participantIdentify string, trackSid string)
	OnTrackSubscriptionFailed(participantIdentify string, trackSid string, err string)
	OnTrackPublished(participantIdentify string, track *RemoteTrack)
	OnTrackUnpublished(participantIdentify string, publicationSid string)
	OnDataPacket(participantIdentify string, value interface{})
	OnTranscriptionReceived(participantIdentify string, trackSid string, segments []*pb_room.TranscriptionSegment)
}
