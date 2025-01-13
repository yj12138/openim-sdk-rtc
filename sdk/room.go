package sdk

import (
	"log"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"

	// pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/pion/webrtc/v4"
)

type Room struct {
	url         string
	token       string
	livekitRoom *lksdk.Room
	listener    RoomListener
}

// RoomCallback
func (r *Room) onDisconnected() {

}
func (r *Room) onDisconnectedWithReason(reason lksdk.DisconnectionReason) {

}
func (r *Room) onParticipantConnected(rp *lksdk.RemoteParticipant) {

}
func (r *Room) onParticipantDisconnected(rp *lksdk.RemoteParticipant) {

}
func (r *Room) onActiveSpeakersChanged(ps []lksdk.Participant) {

}
func (r *Room) onRoomMetadataChanged(metadata string) {

}
func (r *Room) onReconnecting() {

}

func (r *Room) onReconnected() {

}

// ParticipantCallback
func (r *Room) onLocalTrackPublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {

}
func (r *Room) onLocalTrackUnpublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {

}

// for all participants
func (r *Room) onTrackMuted(pub lksdk.TrackPublication, p lksdk.Participant) {

}
func (r *Room) onTrackUnmuted(pub lksdk.TrackPublication, p lksdk.Participant) {

}
func (r *Room) onMetadataChanged(oldMetadata string, p lksdk.Participant) {

}
func (r *Room) onAttributesChanged(changed map[string]string, p lksdk.Participant) {

}
func (r *Room) onIsSpeakingChanged(p lksdk.Participant) {

}
func (r *Room) onConnectionQualityChanged(update *livekit.ConnectionQualityInfo, p lksdk.Participant) {

}

// for remote participants
func (r *Room) onTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("recv track:", rp.Identity(), track.ID())

}

func (r *Room) onTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {

}
func (r *Room) onTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {

}
func (r *Room) onTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {

}
func (r *Room) onTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {

}

// Deprecated: Use OnDataPacket instead
func (r *Room) onDataReceived(data []byte, params lksdk.DataReceiveParams) {

}

func (r *Room) onDataPacket(data lksdk.DataPacket, params lksdk.DataReceiveParams) {

}
func (r *Room) onTranscriptionReceived(transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {

}

func (r *Room) Connect(url, token string, listener RoomListener) {
	r.url = url
	r.token = token
	go func() {
		livekitRoomCallback := &lksdk.RoomCallback{
			OnDisconnected:            r.onDisconnected,
			OnDisconnectedWithReason:  r.onDisconnectedWithReason,
			OnParticipantConnected:    r.onParticipantConnected,
			OnParticipantDisconnected: r.onParticipantDisconnected,
			OnActiveSpeakersChanged:   r.onActiveSpeakersChanged,
			OnRoomMetadataChanged:     r.onRoomMetadataChanged,
			OnReconnecting:            r.onReconnecting,
			OnReconnected:             r.onReconnected,
			ParticipantCallback: lksdk.ParticipantCallback{
				OnLocalTrackPublished:      r.onLocalTrackPublished,
				OnLocalTrackUnpublished:    r.onLocalTrackUnpublished,
				OnTrackMuted:               r.onTrackMuted,
				OnTrackUnmuted:             r.onTrackUnmuted,
				OnMetadataChanged:          r.onMetadataChanged,
				OnAttributesChanged:        r.onAttributesChanged,
				OnIsSpeakingChanged:        r.onIsSpeakingChanged,
				OnConnectionQualityChanged: r.onConnectionQualityChanged,
				OnTrackSubscribed:          r.onTrackSubscribed,
				OnTrackUnsubscribed:        r.onTrackUnsubscribed,
				OnTrackSubscriptionFailed:  r.onTrackSubscriptionFailed,
				OnTrackPublished:           r.onTrackPublished,
				OnTrackUnpublished:         r.onTrackUnpublished,
				OnDataReceived:             r.onDataReceived,
				OnDataPacket:               r.onDataPacket,
				OnTranscriptionReceived:    r.onTranscriptionReceived,
			},
		}
		livekitRoom, err := lksdk.ConnectToRoomWithToken(url, token, livekitRoomCallback, lksdk.WithAutoSubscribe(true))
		if err != nil {
			log.Println(err.Error())
			return
		}
		r.livekitRoom = livekitRoom
	}()
}

func (r *Room) Disconnect() {
	if r.livekitRoom != nil {
		r.livekitRoom.Disconnect()
	}
}

func (r *Room) PublicTrack() {

}

func (r *Room) UnpublishTrack() {
}

func (r *Room) PublishData() {
}

func (r *Room) SetSubscribed() {
}

func (r *Room) UpdateLocalMetadata() {
}

func (r *Room) UpdateLocalName() {
}

func (r *Room) GetSessionStats() {
}

func NewRoom(listener RoomListener) *Room {
	room := &Room{
		listener: listener,
	}
	return room
}
