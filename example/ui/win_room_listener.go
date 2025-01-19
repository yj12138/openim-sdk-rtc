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
	log.Println("Room disconnected:", room.GetRoomName())
}

func (r *RoomListenerImpl) OnDisconnectedWithReason(room *sdk.Room, reason lksdk.DisconnectionReason) {
	log.Println("Room disconnected with reason:", room.GetRoomName(), string(reason))
}

func (r *RoomListenerImpl) OnParticipantConnected(room *sdk.Room, rp *lksdk.RemoteParticipant) {
	log.Println("Participant connected:", room.GetRoomName(), rp.Identity())
}

func (r *RoomListenerImpl) OnParticipantDisconnected(room *sdk.Room, rp *lksdk.RemoteParticipant) {
	log.Println("Participant disconnected:", room.GetRoomName(), rp.Identity())
}

func (r *RoomListenerImpl) OnActiveSpeakersChanged(room *sdk.Room, ps []lksdk.Participant) {
	log.Println("Active speakers changed:", room.GetRoomName())
}

func (r *RoomListenerImpl) OnRoomMetadataChanged(room *sdk.Room, metaData string) {
	log.Println("Room metadata changed:", room.GetRoomName(), metaData)
}

func (r *RoomListenerImpl) OnReconnecting(room *sdk.Room) {
	log.Println("Room reconnecting:", room.GetRoomName())
}

func (r *RoomListenerImpl) OnReconnected(room *sdk.Room) {
	log.Println("Room reconnected:", room.GetRoomName())
}

func (r *RoomListenerImpl) OnLocalTrackPublished(room *sdk.Room, ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println("Local track published:", room.GetRoomName(), ltp.Track().ID(), lp.Identity())
}

func (r *RoomListenerImpl) OnLocalTrackUnpublished(room *sdk.Room, ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println("Local track unpublished:", room.GetRoomName(), ltp.Track().ID(), lp.Identity())
}

func (r *RoomListenerImpl) OnTrackMuted(room *sdk.Room, tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println("Track muted:", room.GetRoomName(), tp.Track().ID(), p.Identity())
}

func (r *RoomListenerImpl) OnTrackUnmuted(room *sdk.Room, tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println("Track unmuted:", room.GetRoomName(), tp.Track().ID(), p.Identity())
}

func (r *RoomListenerImpl) OnMetadataChanged(room *sdk.Room, oldMetaData string, p lksdk.Participant) {
	log.Println("Metadata changed:", room.GetRoomName(), oldMetaData, p.Identity())
}

func (r *RoomListenerImpl) OnAttributesChanged(room *sdk.Room, changed map[string]string, p lksdk.Participant) {
	log.Println("Attributes changed:", room.GetRoomName(), changed, p.Identity())
}

func (r *RoomListenerImpl) OnIsSpeakingChanged(room *sdk.Room, p lksdk.Participant) {
	log.Println("Is speaking changed:", room.GetRoomName(), p.Identity())
}

func (r *RoomListenerImpl) OnConnectionQualityChanged(room *sdk.Room, update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	log.Println("Connection quality changed:", room.GetRoomName(), update.String(), p.Identity())
}

func (r *RoomListenerImpl) OnTrackSubscribed(room *sdk.Room, track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscribed:", room.GetRoomName(), track.ID(), publication.Track().ID(), rp.Identity())
}

func (r *RoomListenerImpl) OnTrackUnsubscribed(room *sdk.Room, track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unsubscribed:", room.GetRoomName(), track.ID(), publication.Track().ID(), rp.Identity())
}

func (r *RoomListenerImpl) OnTrackSubscriptionFailed(room *sdk.Room, sid string, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscription failed:", room.GetRoomName(), sid, rp.Identity())
}

func (r *RoomListenerImpl) OnTrackPublished(room *sdk.Room, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track published:", room.GetRoomName(), publication.Track().ID(), rp.Identity())
}

func (r *RoomListenerImpl) OnTrackUnpublished(room *sdk.Room, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unpublished:", room.GetRoomName(), publication.Track().ID(), rp.Identity())
}

func (r *RoomListenerImpl) OnDataPacket(room *sdk.Room, data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	userDataPackage, ok := data.(*lksdk.UserDataPacket)
	if ok {
		log.Println("Data packet received:", room.GetRoomName(), params.SenderIdentity, string(userDataPackage.Payload))
	}
}

func (r *RoomListenerImpl) OnTranscriptionReceived(room *sdk.Room, transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	log.Println("Transcription received:", room.GetRoomName(), transcriptionSegments, p, publication)
}

func NewRoomListener() sdk.RoomListener {
	return &RoomListenerImpl{}
}
