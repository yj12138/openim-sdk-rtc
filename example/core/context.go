package core

import (
	"log"
	"time"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/example/io"
	"github.com/openimsdk/openim-rtc/sdk"
	"github.com/pion/webrtc/v4"
)

var livekitHost string
var livekitApiKey string
var livekitApiSecret string

var room *sdk.Room
var audioTrack *sdk.Track
var micPhone *io.MicPhone
var speaker *io.Speaker

func Init(host string, apiKey string, apiSecret string) {
	livekitHost = host
	livekitApiKey = apiKey
	livekitApiSecret = apiSecret
	room = sdk.NewRoom(NewRoomCallBack())
	// 每帧时间 = 1 / 采样率（秒）
	audioTrack = sdk.NewAudioTrack("micphone_track", sdk.MimeTypeOpus)
	micPhone = io.NewMicPhone(func(data []byte, frameCount uint32) {
		sampleRate := 44100
		frameDuration := time.Duration(int64(time.Second)/int64(sampleRate)) * time.Duration(frameCount)
		audioTrack.WriteData(data, frameDuration)
	})
	speaker = io.NewSpeaker()
}

func ConnectRoom(roomName string, identify string) {
	room.ConnectBySecret(livekitHost, livekitApiKey, livekitApiSecret, roomName, identify)
}

func OnDisconnected() {
	log.Println("Room disconnected:")
}

func OnDisconnectedWithReason(reason lksdk.DisconnectionReason) {
	log.Println("Room disconnected with reason:", string(reason))
}

func OnParticipantConnected(rp *lksdk.RemoteParticipant) {
	log.Println("Participant connected:", rp.Identity())
}

func OnParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	log.Println("Participant disconnected:", rp.Identity())
}

func OnActiveSpeakersChanged(ps []lksdk.Participant) {
	log.Println("Active speakers changed:")
}

func OnRoomMetadataChanged(metaData string) {
	log.Println("Room metadata changed:", metaData)
}

func OnReconnecting() {
	log.Println("Reconnecting:")
}

func OnReconnected() {
	log.Println("Room reconnected:")
}

func OnLocalTrackPublished(ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println(lp.Identity(), "Local track published:", ltp.Track().ID())
}

func OnLocalTrackUnpublished(ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println(lp.Identity(), "Local track unpublished:", ltp.Track().ID())
}

func OnTrackMuted(tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println(p.Identity(), "Track muted:", tp.Track().ID())
}

func OnTrackUnmuted(tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println(p.Identity(), "Track unmuted:", tp.Track().ID())
}

func OnMetadataChanged(oldMetaData string, p lksdk.Participant) {
	log.Println(p.Identity(), "Metadata changed:", oldMetaData)
}

func OnAttributesChanged(changed map[string]string, p lksdk.Participant) {
	log.Println(p.Identity(), "Attributes changed:", changed)
}

func OnIsSpeakingChanged(p lksdk.Participant) {
	log.Println(p.Identity(), "Is speaking changed:")
}

func OnConnectionQualityChanged(update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	log.Println(p.Identity(), "Connection quality changed:", update.String())
}

func OnTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscribed:", track.ID(), publication.Track().ID(), rp.Identity())
	go func() {
		log.Println(track.Codec().MimeType)
		for {
			pkt, _, err := track.ReadRTP()
			if err != nil {
				break
			}
			data := pkt.Payload
			speaker.WriteData(data)
		}
	}()
}

func OnTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unsubscribed:", track.ID(), publication.Track().ID(), rp.Identity())
}

func OnTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscription failed:", sid, rp.Identity())
}

func OnTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("OnTrackPublished", publication.MimeType(), rp.Identity())
}

func OnTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unpublished:", publication.Track().ID(), rp.Identity())
}

func OnDataPacket(data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	userDataPackage, ok := data.(*lksdk.UserDataPacket)
	if ok {
		log.Println("Data packet received:", params.SenderIdentity, string(userDataPackage.Payload))
	}
}

func OnTranscriptionReceived(transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	log.Println("Transcription received:", transcriptionSegments, p, publication)
}

func NewRoomCallBack() *lksdk.RoomCallback {
	return &lksdk.RoomCallback{
		OnDisconnected:            OnDisconnected,
		OnDisconnectedWithReason:  OnDisconnectedWithReason,
		OnParticipantConnected:    OnParticipantConnected,
		OnParticipantDisconnected: OnParticipantDisconnected,
		OnActiveSpeakersChanged:   OnActiveSpeakersChanged,
		OnRoomMetadataChanged:     OnRoomMetadataChanged,
		OnReconnecting:            OnReconnecting,
		OnReconnected:             OnReconnected,
		ParticipantCallback: lksdk.ParticipantCallback{
			OnLocalTrackPublished:   OnLocalTrackPublished,
			OnLocalTrackUnpublished: OnLocalTrackUnpublished,
			// for all participants
			OnTrackMuted:               OnTrackMuted,
			OnTrackUnmuted:             OnTrackUnmuted,
			OnMetadataChanged:          OnMetadataChanged,
			OnAttributesChanged:        OnAttributesChanged,
			OnIsSpeakingChanged:        OnIsSpeakingChanged,
			OnConnectionQualityChanged: OnConnectionQualityChanged,

			// for remote participants
			OnTrackSubscribed:         OnTrackSubscribed,
			OnTrackUnsubscribed:       OnTrackUnsubscribed,
			OnTrackSubscriptionFailed: OnTrackSubscriptionFailed,
			OnTrackPublished:          OnTrackPublished,
			OnTrackUnpublished:        OnTrackUnpublished,
			OnDataPacket:              OnDataPacket,
			OnTranscriptionReceived:   OnTranscriptionReceived,
		},
	}
}
