package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"

	"github.com/livekit/protocol/livekit"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	"log"
)

type Room struct {
	host        string
	token       string
	roomName    string
	owner       string
	livekitRoom *lksdk.Room

	listener OnRoomListener
}

func (r *Room) createCallBack() *lksdk.RoomCallback {
	callBack := &lksdk.RoomCallback{
		// room
		OnDisconnected:            r.onDisconnected,
		OnDisconnectedWithReason:  r.onDisconnectedWithReason,
		OnParticipantConnected:    r.onParticipantConnected,
		OnParticipantDisconnected: r.onParticipantDisconnected,
		OnActiveSpeakersChanged:   r.onActiveSpeakersChanged,
		OnRoomMetadataChanged:     r.onRoomMetadataChanged,
		OnReconnecting:            r.onReconnecting,
		OnReconnected:             r.onReconnected,
		ParticipantCallback: lksdk.ParticipantCallback{
			OnLocalTrackPublished:   r.onLocalTrackPublished,
			OnLocalTrackUnpublished: r.onLocalTrackUnpublished,

			// for all participants
			OnTrackMuted:               r.onTrackMuted,
			OnTrackUnmuted:             r.onTrackUnmuted,
			OnMetadataChanged:          r.onMetadataChanged,
			OnAttributesChanged:        r.onAttributesChanged,
			OnIsSpeakingChanged:        r.onIsSpeakingChanged,
			OnConnectionQualityChanged: r.onConnectionQualityChanged,

			// for remote participants
			OnTrackSubscribed:         r.onTrackSubscribed,
			OnTrackUnsubscribed:       r.onTrackUnsubscribed,
			OnTrackSubscriptionFailed: r.onTrackSubscriptionFailed,
			OnTrackPublished:          r.onTrackPublished,
			OnTrackUnpublished:        r.onTrackUnpublished,
			OnDataReceived:            func(data []byte, params lksdk.DataReceiveParams) {},
			OnDataPacket:              r.onDataPacket,
			OnTranscriptionReceived:   r.onTranscriptionReceived,
		},
		// participant

	}

	return callBack
}

func (r *Room) ConnectByToken(host, token string) {
	r.host = host
	r.token = token
	livekitRoom, err := lksdk.ConnectToRoomWithToken(host, token, r.createCallBack(), lksdk.WithAutoSubscribe(true))
	if err != nil {
		log.Println(err.Error())
		return
	}
	r.livekitRoom = livekitRoom
}

func (r *Room) ConnectBySecret(host, apiKey, apiSecret, roomName, identify string) {
	r.roomName = roomName
	r.owner = identify
	room, err := lksdk.ConnectToRoom(host, lksdk.ConnectInfo{
		APIKey:              apiKey,
		APISecret:           apiSecret,
		RoomName:            roomName,
		ParticipantIdentity: identify,
	}, r.createCallBack(), lksdk.WithAutoSubscribe(true))
	if err != nil {
		panic(err)
	}
	r.livekitRoom = room
}

func (r *Room) Disconnect() {
	if r.livekitRoom != nil {
		r.livekitRoom.Disconnect()
	}
	r.livekitRoom = nil
}

func (r *Room) IsConnSuc() bool {
	return r.livekitRoom != nil
}

func (r *Room) checkConn() bool {
	if r.IsConnSuc() {
		return true
	} else {
		log.Panic("not connect to room")
		return false
	}
}

func (r *Room) GetRoomName() string {
	return r.roomName
}

func (r *Room) GetOwner() string {
	return r.owner
}

func (r *Room) GetLocalParticipant() *lksdk.LocalParticipant {
	return r.livekitRoom.LocalParticipant
}

func (r *Room) GetAllParticipantId() []string {
	res := make([]string, 0)
	if r.checkConn() {
		rps := r.livekitRoom.GetRemoteParticipants()
		for _, rp := range rps {
			res = append(res, rp.Identity())
		}
	}
	return res
}

// listener
func (r *Room) onDisconnected() {
	r.listener.OnDisconnectedWithReason(pb_participant.DisconnectReason_UNKNOWN_REASON)
}
func (r *Room) onDisconnectedWithReason(reason lksdk.DisconnectionReason) {
	res := pb_participant.DisconnectReason_UNKNOWN_REASON
	if reason == lksdk.LeaveRequested {
		res = pb_participant.DisconnectReason_ROOM_CLOSED
	} else if reason == lksdk.UserUnavailable {
		res = pb_participant.DisconnectReason_USER_UNAVAILABLE
	} else if reason == lksdk.RejectedByUser {
		res = pb_participant.DisconnectReason_USER_REJECTED
	} else if reason == lksdk.Failed {
		res = pb_participant.DisconnectReason_JOIN_FAILURE
	}
	r.listener.OnDisconnectedWithReason(res)
}
func (r *Room) onParticipantConnected(rp *lksdk.RemoteParticipant) {
	r.listener.OnParticipantConnected()
}
func (r *Room) onParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	r.listener.OnParticipantDisconnected()
}
func (r *Room) onActiveSpeakersChanged(ps []lksdk.Participant) {
	r.listener.OnActiveSpeakersChanged()
}
func (r *Room) onRoomMetadataChanged(metadata string) {
	r.listener.OnRoomMetadataChanged(metadata)
}
func (r *Room) onReconnecting() {
	r.listener.OnReconnecting()
}
func (r *Room) onReconnected() {
	r.listener.OnReconnected()
}

func (r *Room) onLocalTrackPublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	r.listener.OnLocalTrackPublished()
}
func (r *Room) onLocalTrackUnpublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	r.listener.OnLocalTrackUnpublished()
}
func (r *Room) onTrackMuted(pub lksdk.TrackPublication, p lksdk.Participant) {
	r.listener.OnTrackMuted()
}
func (r *Room) onTrackUnmuted(pub lksdk.TrackPublication, p lksdk.Participant) {
	r.listener.OnTrackUnmuted()
}
func (r *Room) onMetadataChanged(oldMetadata string, p lksdk.Participant) {
	r.listener.OnMetadataChanged()
}
func (r *Room) onAttributesChanged(changed map[string]string, p lksdk.Participant) {
	r.listener.OnAttributesChanged()
}

func (r *Room) onIsSpeakingChanged(p lksdk.Participant) {
	r.listener.OnIsSpeakingChanged()
}

func (r *Room) onConnectionQualityChanged(update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	r.listener.OnConnectionQualityChanged()
}

func (r *Room) onTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackSubscribed()
}
func (r *Room) onTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackUnsubscribed()
}
func (r *Room) onTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackSubscriptionFailed()
}
func (r *Room) onTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackPublished()
}
func (r *Room) onTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackUnpublished()
}
func (r *Room) onDataPacket(data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	r.listener.OnDataPacket()
}
func (r *Room) onTranscriptionReceived(transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	r.listener.OnTranscriptionReceived()
}

func NewRoom(listener OnRoomListener) *Room {
	r := &Room{
		listener: listener,
	}
	return r
}
