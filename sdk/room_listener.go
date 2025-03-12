package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/pion/webrtc/v4"
)

type OnRoomListener interface {
	// room
	OnDisconnectedWithReason(pb_participant.DisconnectReason)
	OnParticipantConnected(*lksdk.RemoteParticipant)
	OnParticipantDisconnected(*lksdk.RemoteParticipant)
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
	OnParticipantNameChanged(participantIdentify string, newName string, oldName string)
	OnParticipantMetadataChanged(participantIdentify string, metadata string)
	OnParticipantAttributesChanged(participantIdentify string, changed map[string]string)
	OnIsSpeakingChanged(participantIdentify []string)
	OnConnectionQualityChanged(participantIdentify string, quality pb_room.ConnectionQuality)

	// for remote participants
	OnTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnTrackSubscriptionFailed(participantIdentify string, trackSid string, err string)
	OnTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnDataPacket(participantIdentify string, packet lksdk.DataPacket)
	OnTranscriptionReceived(participantIdentify string, trackSid string, segments []*lksdk.TranscriptionSegment)
}
