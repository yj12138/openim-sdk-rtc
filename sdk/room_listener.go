package sdk

import (
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
)

type OnRoomListener interface {
	// room
	OnDisconnectedWithReason(pb_participant.DisconnectReason)
	OnParticipantConnected()
	OnParticipantDisconnected()
	OnActiveSpeakersChanged()
	OnRoomMetadataChanged(metadata string)
	OnReconnecting()
	OnReconnected()
	// for local participants
	OnLocalTrackPublished()
	OnLocalTrackUnpublished()

	// for all participants
	OnTrackMuted()
	OnTrackUnmuted()
	OnMetadataChanged()
	OnAttributesChanged()
	OnIsSpeakingChanged()
	OnConnectionQualityChanged()

	// for remote participants
	OnTrackSubscribed()
	OnTrackUnsubscribed()
	OnTrackSubscriptionFailed()
	OnTrackPublished()
	OnTrackUnpublished()
	OnDataPacket()
	OnTranscriptionReceived()
}
