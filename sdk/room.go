package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"

	"log"

	"github.com/livekit/protocol/livekit"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type Room struct {
	*lksdk.Room
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
	}

	return callBack
}

func (r *Room) GetConnectState() pb_room.ConnectionState {
	state := r.ConnectionState()
	if state == lksdk.ConnectionStateConnected {
		return pb_room.ConnectionState_CONN_CONNECTED
	} else if state == lksdk.ConnectionStateDisconnected {
		return pb_room.ConnectionState_CONN_DISCONNECTED
	} else if state == lksdk.ConnectionStateReconnecting {
		return pb_room.ConnectionState_CONN_RECONNECTING
	}
	return pb_room.ConnectionState_CONN_DISCONNECTED
}

func (r *Room) GetAllParticipantIdentifys() []string {
	res := make([]string, 0)
	rps := r.GetRemoteParticipants()
	for _, rp := range rps {
		res = append(res, rp.Identity())
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
	r.listener.OnParticipantConnected(rp)
}
func (r *Room) onParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	r.listener.OnParticipantDisconnected(rp)
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
	quality := pb_room.ConnectionQuality_QUALITY_POOR
	if update.Quality == livekit.ConnectionQuality_EXCELLENT {
		quality = pb_room.ConnectionQuality_QUALITY_EXCELLENT
	} else if update.Quality == livekit.ConnectionQuality_GOOD {
		quality = pb_room.ConnectionQuality_QUALITY_GOOD
	} else if update.Quality == livekit.ConnectionQuality_LOST {
		quality = pb_room.ConnectionQuality_QUALITY_LOST
	}
	r.listener.OnConnectionQualityChanged(p.Identity(), quality)
}

func (r *Room) onTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackSubscribed(rp.Identity(), publication)
}
func (r *Room) onTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackUnsubscribed(rp.Identity(), publication.SID())
}
func (r *Room) onTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackSubscriptionFailed(rp.Identity(), sid, "TODO")
}
func (r *Room) onTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackPublished(rp.Identity(), publication)
}
func (r *Room) onTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	r.listener.OnTrackUnpublished(rp.Identity(), publication.SID())
}
func (r *Room) onDataPacket(data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	packet := data.ToProto()
	var value interface{} = nil
	if user, ok := packet.Value.(*livekit.DataPacket_User); ok {
		value = &pb_room.UserPacket{
			Topic: *user.User.Topic,
			// TODO
			// Data:  user.User.Payload,
		}
	}
	if sipDtmf, ok := packet.Value.(*livekit.DataPacket_SipDtmf); ok {
		value = &pb_room.SipDTMF{
			Digit: sipDtmf.SipDtmf.Digit,
			Code:  sipDtmf.SipDtmf.Code,
		}
	}
	if value != nil {
		r.listener.OnDataPacket(packet.ParticipantIdentity, nil)
	}
}
func (r *Room) onTranscriptionReceived(transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	segments := make([]*pb_room.TranscriptionSegment, len(transcriptionSegments))
	for i, segment := range transcriptionSegments {
		segments[i] = &pb_room.TranscriptionSegment{
			Id:        segment.ID,
			Text:      segment.Text,
			StartTime: segment.StartTime,
			EndTime:   segment.EndTime,
			Final:     segment.Final,
			Language:  segment.Language,
		}
	}

	r.listener.OnTranscriptionReceived(p.Identity(), publication.SID(), segments)
}

func ConnectByToken(host, token string, listener OnRoomListener) *Room {
	room := &Room{}
	room.listener = listener
	livekitRoom, err := lksdk.ConnectToRoomWithToken(host, token, room.createCallBack(), lksdk.WithAutoSubscribe(true))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	room.Room = livekitRoom
	return room
}

func ConnectBySecret(host, apiKey, apiSecret, roomName, identify string, listener OnRoomListener) *Room {
	room := &Room{}
	room.listener = listener
	livekitRoom, err := lksdk.ConnectToRoom(host, lksdk.ConnectInfo{
		APIKey:              apiKey,
		APISecret:           apiSecret,
		RoomName:            roomName,
		ParticipantIdentity: identify,
	}, room.createCallBack(), lksdk.WithAutoSubscribe(true))
	if err != nil {
		panic(err)
	}
	room.Room = livekitRoom
	return room
}
