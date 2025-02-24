package sdk

type OnRoomListener interface {
	OnDisconnected()
	OnDisconnectedWithReason(reason string)
	OnParticipantConnected()
	OnParticipantDisconnected()
	OnActiveSpeakersChanged()
	OnRoomMetadataChanged(metadata string)
	OnReconnecting()
	OnReconnected()
}
