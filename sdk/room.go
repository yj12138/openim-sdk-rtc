package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"

	"github.com/livekit/protocol/livekit"
	pb_common "github.com/openimsdk/openim-rtc/proto/go/common"
	"log"
)

type Room struct {
	host        string
	token       string
	roomName    string
	owner       string
	livekitRoom *lksdk.Room

	listener OnRoomListener

	remoteParticipants map[string]*RemoteParticipant
	remoteTracks       map[string]*RemoteTrack
}

func NewRoom(listener OnRoomListener) *Room {
	r := &Room{
		listener:           listener,
		remoteParticipants: make(map[string]*RemoteParticipant),
		remoteTracks:       make(map[string]*RemoteTrack),
	}
	return r
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
			OnTrackMuted:               r.onParticipantTrackMuted,
			OnTrackUnmuted:             r.onParticipantTrackUnmuted,
			OnMetadataChanged:          r.onParticipantMetadataChanged,
			OnAttributesChanged:        r.onParticipantAttributesChanged,
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
	r.listener.OnDisconnectedWithReason(pb_common.DisconnectReason_UNKNOWN_REASON)
}
func (r *Room) onDisconnectedWithReason(reason lksdk.DisconnectionReason) {
	res := pb_common.DisconnectReason_UNKNOWN_REASON
	if reason == lksdk.LeaveRequested {
		res = pb_common.DisconnectReason_ROOM_CLOSED
	} else if reason == lksdk.UserUnavailable {
		res = pb_common.DisconnectReason_USER_UNAVAILABLE
	} else if reason == lksdk.RejectedByUser {
		res = pb_common.DisconnectReason_USER_REJECTED
	} else if reason == lksdk.Failed {
		res = pb_common.DisconnectReason_JOIN_FAILURE
	}
	r.listener.OnDisconnectedWithReason(res)
}
func (r *Room) onParticipantConnected(rp *lksdk.RemoteParticipant) {
	remoteParticipant := NewRemoteParticipant(rp)
	r.remoteParticipants[remoteParticipant.LiveKitRemoteParticipant.Identity()] = remoteParticipant
	r.listener.OnParticipantConnected(remoteParticipant)
}
func (r *Room) onParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	remoteParticipant, ok := r.remoteParticipants[rp.Identity()]
	if ok {
		r.listener.OnParticipantDisconnected(remoteParticipant)
		delete(r.remoteParticipants, rp.Identity())
	} else {
		log.Printf("not find participant %s", rp.Identity())
	}
}
func (r *Room) onActiveSpeakersChanged(ps []lksdk.Participant) {
	participantIdentities := make([]string, len(ps))
	for i, participant := range ps {
		participantIdentities[i] = participant.Identity()
	}
	r.listener.OnActiveSpeakersChanged(participantIdentities)
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
	r.listener.OnLocalTrackPublished(publication.SID())
}
func (r *Room) onLocalTrackUnpublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	r.listener.OnLocalTrackUnpublished(publication.SID())
}
func (r *Room) onParticipantTrackMuted(pub lksdk.TrackPublication, p lksdk.Participant) {
	r.listener.OnParticipantTrackMuted(p.Identity(), pub.SID())
}
func (r *Room) onParticipantTrackUnmuted(pub lksdk.TrackPublication, p lksdk.Participant) {
	r.listener.OnParticipantTrackUnmuted(p.Identity(), pub.SID())
}
func (r *Room) onParticipantMetadataChanged(oldMetadata string, p lksdk.Participant) {
	r.listener.OnParticipantMetadataChanged(p.Identity(), p.Metadata())
}
func (r *Room) onParticipantAttributesChanged(changed map[string]string, p lksdk.Participant) {
	r.listener.OnParticipantAttributesChanged(p.Identity(), changed)
}

func (r *Room) onIsSpeakingChanged(p lksdk.Participant) {
	ps := make([]string, 0)
	ps = append(ps, p.Identity())
	r.listener.OnIsSpeakingChanged(ps)
}

func (r *Room) onConnectionQualityChanged(update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	quality := pb_common.ConnectionQuality_QUALITY_POOR
	if update.Quality == livekit.ConnectionQuality_EXCELLENT {
		quality = pb_common.ConnectionQuality_QUALITY_EXCELLENT
	} else if update.Quality == livekit.ConnectionQuality_GOOD {
		quality = pb_common.ConnectionQuality_QUALITY_GOOD
	} else if update.Quality == livekit.ConnectionQuality_LOST {
		quality = pb_common.ConnectionQuality_QUALITY_LOST
	}
	r.listener.OnConnectionQualityChanged(p.Identity(), quality)
}

func (r *Room) onTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	remoteTrack := NewRemoteTrack(track, publication, rp)
	r.remoteTracks[publication.SID()] = remoteTrack
	r.listener.OnTrackSubscribed(remoteTrack)
}
func (r *Room) onTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	remoteTrack, ok := r.remoteTracks[publication.SID()]
	if ok {
		r.listener.OnTrackUnsubscribed(remoteTrack)
		delete(r.remoteTracks, publication.SID())
	} else {
		log.Printf("not find RemoteTrack %s", publication.SID())
	}
}
func (r *Room) onTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackSubscriptionFailed(rp.Identity(), sid, "TODO")
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
