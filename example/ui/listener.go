package ui

import (
	"log"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/pion/webrtc/v4"
)

type RoomListener struct {
	RoomHandle uint64
}

// for room
func (l *RoomListener) OnDisconnectedWithReason(reason pb_participant.DisconnectReason) {
	log.Println("OnDisconnectedWithReason", reason)
}
func (l *RoomListener) OnParticipantConnected(rp *lksdk.RemoteParticipant) {
	log.Println("OnParticipantConnected", rp.Identity())
}
func (l *RoomListener) OnParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	log.Println("OnParticipantDisconnected", rp.Identity())
}
func (l *RoomListener) OnActiveSpeakersChanged(participantIdentities []string) {
	log.Println("OnActiveSpeakersChanged", participantIdentities)
}

func (l *RoomListener) OnRoomMetadataChanged(metadata string) {
	log.Println("OnRoomMetadataChanged", metadata)
}
func (l *RoomListener) OnReconnecting() {
	log.Println("OnReconnecting")
}
func (l *RoomListener) OnReconnected() {
	log.Println("OnReconnected")
}

// for local participants
func (l *RoomListener) OnLocalTrackPublished(trackSid string) {
	log.Println("OnLocalTrackPublished", trackSid)
}
func (l *RoomListener) OnLocalTrackUnpublished(publicationSid string) {
	log.Println("OnLocalTrackUnpublished", publicationSid)
}

// for all participants
func (l *RoomListener) OnParticipantTrackMuted(participantIdentify string, trackSid string) {
	log.Println("OnParticipantTrackMuted", participantIdentify, trackSid)
}
func (l *RoomListener) OnParticipantTrackUnmuted(participantIdentify string, trackSid string) {
	log.Println("OnParticipantTrackUnmuted", participantIdentify, trackSid)
}
func (l *RoomListener) OnParticipantNameChanged(participantIdentify string, newName string, oldName string) {
	log.Println("OnParticipantNameChanged", participantIdentify, newName, oldName)
}
func (l *RoomListener) OnParticipantMetadataChanged(participantIdentify string, metadata string) {
	log.Println("OnParticipantMetadataChanged", participantIdentify, metadata)
}
func (l *RoomListener) OnParticipantAttributesChanged(participantIdentify string, changed map[string]string) {
	log.Println("OnParticipantAttributesChanged", participantIdentify, changed)
}
func (l *RoomListener) OnIsSpeakingChanged(participantIdentities []string) {
	log.Println("OnIsSpeakingChanged", participantIdentities)
}
func (l *RoomListener) OnConnectionQualityChanged(participantIdentify string, quality pb_room.ConnectionQuality) {
	log.Println("OnConnectionQualityChanged", participantIdentify, quality)
}

// for remote participants
func (l *RoomListener) OnTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("OnTrackSubscribed", rp.Identity(), publication.SID(), track.Codec().Channels, track.Codec().MimeType, track.Codec().PayloadType)
	go func() {
		for {
			rtp, _, err := track.ReadRTP()
			if err != nil {
				break
			}
			// log.Println("Recv Frame ", track.Codec(), len(rtp.Payload))
			context.Speaker.Write(rtp.Payload)
		}
	}()
}
func (l *RoomListener) OnTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("OnTrackUnsubscribed", rp.Identity(), publication.SID())
}
func (l *RoomListener) OnTrackSubscriptionFailed(participantIdentify string, trackSid string, err string) {
	log.Println("OnTrackSubscriptionFailed", participantIdentify, trackSid, err)
}

func (l *RoomListener) OnTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("OnTrackPublished", rp.Identity(), publication.SID())
}

func (l *RoomListener) OnTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("OnTrackUnpublished", rp.Identity(), publication.SID())
}

func (l *RoomListener) OnDataPacket(participantIdentify string, _packet lksdk.DataPacket) {
	packet := _packet.ToProto()
	if user, ok := packet.Value.(*livekit.DataPacket_User); ok {
		log.Println("Recv User Data", participantIdentify, user)
	}

	if sipDtmf, ok := packet.Value.(*livekit.DataPacket_SipDtmf); ok {
		log.Println("Recv SipDtmf Data", participantIdentify, sipDtmf)
	}
}

func (l *RoomListener) OnTranscriptionReceived(participantIdentify string, trackSid string, segments []*lksdk.TranscriptionSegment) {
	log.Println("OnTranscriptionReceived", participantIdentify, trackSid, segments)
}

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}
