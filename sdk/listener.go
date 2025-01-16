package sdk

import (
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
)

type RoomListener interface {
	// RoomCallback
	OnDisconnected(room *Room)
	OnDisconnectedWithReason(room *Room, reason lksdk.DisconnectionReason)
	OnParticipantConnected(room *Room, rp *lksdk.RemoteParticipant)
	OnParticipantDisconnected(room *Room, rp *lksdk.RemoteParticipant)
	OnActiveSpeakersChanged(room *Room, ps []lksdk.Participant)
	OnRoomMetadataChanged(room *Room, metaData string)
	OnReconnecting(room *Room)
	OnReconnected(room *Room)
	// ParticipantCallback
	OnLocalTrackPublished(room *Room, ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant)
	OnLocalTrackUnpublished(room *Room, ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant)
	// for all participants
	OnTrackMuted(room *Room, tp lksdk.TrackPublication, p lksdk.Participant)
	OnTrackUnmuted(room *Room, tp lksdk.TrackPublication, p lksdk.Participant)
	OnMetadataChanged(room *Room, oldMetaData string, p lksdk.Participant)
	OnAttributesChanged(room *Room, changed map[string]string, p lksdk.Participant)
	OnIsSpeakingChanged(room *Room, p lksdk.Participant)
	OnConnectionQualityChanged(room *Room, update *livekit.ConnectionQualityInfo, p lksdk.Participant)
	// for remote participants
	OnTrackSubscribed(room *Room, track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnTrackUnsubscribed(room *Room, track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnTrackSubscriptionFailed(room *Room, sid string, rp *lksdk.RemoteParticipant)
	OnTrackPublished(room *Room, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnTrackUnpublished(room *Room, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant)
	OnDataPacket(room *Room, data lksdk.DataPacket, params lksdk.DataReceiveParams)
	OnTranscriptionReceived(room *Room, transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication)
}
