package ui

import (
	"log"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/sdk"
	"github.com/pion/webrtc/v4"
)

type RoomListenerImpl struct {
}

func (r *RoomListenerImpl) OnDisconnected(room *sdk.Room) {
	log.Println("Room disconnected:", room)
}

func (r *RoomListenerImpl) OnDisconnectedWithReason(room *sdk.Room, reason lksdk.DisconnectionReason) {
	log.Println("Room disconnected with reason:", room, reason)
}

func (r *RoomListenerImpl) OnParticipantConnected(room *sdk.Room, rp *lksdk.RemoteParticipant) {
	log.Println("Participant connected:", room, rp)
}

func (r *RoomListenerImpl) OnParticipantDisconnected(room *sdk.Room, rp *lksdk.RemoteParticipant) {
	log.Println("Participant disconnected:", room, rp)
}

func (r *RoomListenerImpl) OnActiveSpeakersChanged(room *sdk.Room, ps []lksdk.Participant) {
	log.Println("Active speakers changed:", room, ps)
}

func (r *RoomListenerImpl) OnRoomMetadataChanged(room *sdk.Room, metaData string) {
	log.Println("Room metadata changed:", room, metaData)
}

func (r *RoomListenerImpl) OnReconnecting(room *sdk.Room) {
	log.Println("Room reconnecting:", room)
}

func (r *RoomListenerImpl) OnReconnected(room *sdk.Room) {
	log.Println("Room reconnected:", room)
}

func (r *RoomListenerImpl) OnLocalTrackPublished(room *sdk.Room, ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println("Local track published:", room, ltp, lp)
}

func (r *RoomListenerImpl) OnLocalTrackUnpublished(room *sdk.Room, ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println("Local track unpublished:", room, ltp, lp)
}

func (r *RoomListenerImpl) OnTrackMuted(room *sdk.Room, tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println("Track muted:", room, tp, p)
}

func (r *RoomListenerImpl) OnTrackUnmuted(room *sdk.Room, tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println("Track unmuted:", room, tp, p)
}

func (r *RoomListenerImpl) OnMetadataChanged(room *sdk.Room, oldMetaData string, p lksdk.Participant) {
	log.Println("Metadata changed:", room, oldMetaData, p)
}

func (r *RoomListenerImpl) OnAttributesChanged(room *sdk.Room, changed map[string]string, p lksdk.Participant) {
	log.Println("Attributes changed:", room, changed, p)
}

func (r *RoomListenerImpl) OnIsSpeakingChanged(room *sdk.Room, p lksdk.Participant) {
	log.Println("Is speaking changed:", room, p)
}

func (r *RoomListenerImpl) OnConnectionQualityChanged(room *sdk.Room, update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	log.Println("Connection quality changed:", room, update, p)
}

func (r *RoomListenerImpl) OnTrackSubscribed(room *sdk.Room, track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscribed:", room, track, publication, rp)
}

func (r *RoomListenerImpl) OnTrackUnsubscribed(room *sdk.Room, track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unsubscribed:", room, track, publication, rp)
}

func (r *RoomListenerImpl) OnTrackSubscriptionFailed(room *sdk.Room, sid string, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscription failed:", room, sid, rp)
}

func (r *RoomListenerImpl) OnTrackPublished(room *sdk.Room, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track published:", room, publication, rp)
}

func (r *RoomListenerImpl) OnTrackUnpublished(room *sdk.Room, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unpublished:", room, publication, rp)
}

func (r *RoomListenerImpl) OnDataPacket(room *sdk.Room, data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	log.Println("Data packet received:", room, data, params)
}

func (r *RoomListenerImpl) OnTranscriptionReceived(room *sdk.Room, transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	log.Println("Transcription received:", room, transcriptionSegments, p, publication)
}

func NewRoomListener() sdk.RoomListener {
	return &RoomListenerImpl{}
}
