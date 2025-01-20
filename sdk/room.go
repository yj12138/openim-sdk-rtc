package sdk

import (
	"log"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"

	// pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/pion/webrtc/v4"
)

type Room struct {
	host        string
	token       string
	roomName    string
	owner       string
	livekitRoom *lksdk.Room
}

// RoomCallback
func (r *Room) onDisconnected() {
	log.Println(r.GetRoomName())
	listener.OnDisconnected(r)
}
func (r *Room) onDisconnectedWithReason(reason lksdk.DisconnectionReason) {
	listener.OnDisconnectedWithReason(r, reason)
}
func (r *Room) onParticipantConnected(rp *lksdk.RemoteParticipant) {
	listener.OnParticipantConnected(r, rp)
}
func (r *Room) onParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	listener.OnParticipantDisconnected(r, rp)
}
func (r *Room) onActiveSpeakersChanged(ps []lksdk.Participant) {
	listener.OnActiveSpeakersChanged(r, ps)
}
func (r *Room) onRoomMetadataChanged(metadata string) {
	listener.OnRoomMetadataChanged(r, metadata)
}
func (r *Room) onReconnecting() {
	listener.OnReconnecting(r)
}
func (r *Room) onReconnected() {
	listener.OnReconnected(r)
}

// ParticipantCallback
func (r *Room) onLocalTrackPublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	listener.OnLocalTrackPublished(r, publication, lp)
}
func (r *Room) onLocalTrackUnpublished(publication *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	listener.OnLocalTrackUnpublished(r, publication, lp)
}

// for all participants
func (r *Room) onTrackMuted(pub lksdk.TrackPublication, p lksdk.Participant) {
	listener.OnTrackMuted(r, pub, p)
}
func (r *Room) onTrackUnmuted(pub lksdk.TrackPublication, p lksdk.Participant) {
	listener.OnTrackUnmuted(r, pub, p)
}
func (r *Room) onMetadataChanged(oldMetadata string, p lksdk.Participant) {
	listener.OnMetadataChanged(r, oldMetadata, p)
}
func (r *Room) onAttributesChanged(changed map[string]string, p lksdk.Participant) {
	listener.OnAttributesChanged(r, changed, p)
}
func (r *Room) onIsSpeakingChanged(p lksdk.Participant) {
	listener.OnIsSpeakingChanged(r, p)
}
func (r *Room) onConnectionQualityChanged(update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	listener.OnConnectionQualityChanged(r, update, p)
}

// for remote participants
func (r *Room) onTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	listener.OnTrackSubscribed(r, track, publication, rp)
}
func (r *Room) onTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	listener.OnTrackUnsubscribed(r, track, publication, rp)
}
func (r *Room) onTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	listener.OnTrackSubscriptionFailed(r, sid, rp)
}
func (r *Room) onTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	listener.OnTrackPublished(r, publication, rp)
}
func (r *Room) onTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	listener.OnTrackUnpublished(r, publication, rp)
}
func (r *Room) onDataPacket(data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	listener.OnDataPacket(r, data, params)
}
func (r *Room) onTranscriptionReceived(transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	listener.OnTranscriptionReceived(r, transcriptionSegments, p, publication)
}

func (r *Room) createRoomCallBack() *lksdk.RoomCallback {
	return &lksdk.RoomCallback{
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
			OnDataPacket:               r.onDataPacket,
			OnTranscriptionReceived:    r.onTranscriptionReceived,
		},
	}
}

func (r *Room) ConnectByToken(host, token, roomName, identify string) {
	r.host = host
	r.token = token
	go func() {
		livekitRoom, err := lksdk.ConnectToRoomWithToken(host, token, r.createRoomCallBack(), lksdk.WithAutoSubscribe(true))
		if err != nil {
			log.Println(err.Error())
			return
		}
		r.livekitRoom = livekitRoom
	}()
}

func (r *Room) ConnectBySecret(host, apiKey, apiSecret, roomName, identify string) {
	r.roomName = roomName
	r.owner = identify
	room, err := lksdk.ConnectToRoom(host, lksdk.ConnectInfo{
		APIKey:              apiKey,
		APISecret:           apiSecret,
		RoomName:            roomName,
		ParticipantIdentity: identify,
	}, r.createRoomCallBack())
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

func (r *Room) PublicTrack(track *Track) {
	if r.checkConn() {
		_, err := r.livekitRoom.LocalParticipant.PublishTrack(track.liveKitTruck, &lksdk.TrackPublicationOptions{
			VideoWidth:  100,
			VideoHeight: 100,
			Name:        "",
		})
		if err != nil {
			log.Panic(err)
		}
	}
}

func (r *Room) UnpublishTrack() {
}

func (r *Room) checkConn() bool {
	if r.IsConnSuc() {
		return true
	} else {
		log.Panic("not connect to room")
		return false
	}
}

func (r *Room) PublishData(data string) {
	if r.checkConn() {
		r.livekitRoom.LocalParticipant.PublishDataPacket(lksdk.UserData([]byte(data)), lksdk.WithDataPublishReliable(true))
	}
}

func (r *Room) SetSubscribed() {
}

func (r *Room) UpdateLocalMetadata() {
}

func (r *Room) UpdateLocalName() {
}

func (r *Room) GetSessionStats() {
}

func (r *Room) GetRoomName() string {
	return r.roomName
}

func (r *Room) GetOwner() string {
	return r.owner
}

func NewRoom() *Room {
	room := &Room{}
	return room
}
