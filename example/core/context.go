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

func OpenAudioTrack() {
	room.PublicTrack(audioTrack)
}

func StopAudioTrack() {
	room.UnpublishTrack(audioTrack)
}

func SendData(data string) {
	room.PublishData(data)
}

func OpenMicPhone() {
	micPhone.Start()
}

func StopMicPhone() {
	micPhone.Stop()
}

func OpenSpeaker() {
	speaker.Start()
}

func StopSpeakder() {
	speaker.Stop()
}

func onDisconnected() {
	log.Println("Room disconnected:")
}

func onDisconnectedWithReason(reason lksdk.DisconnectionReason) {
	log.Println("Room disconnected with reason:", string(reason))
}

func onParticipantConnected(rp *lksdk.RemoteParticipant) {
	log.Println("Participant connected:", rp.Identity())
}

func onParticipantDisconnected(rp *lksdk.RemoteParticipant) {
	log.Println("Participant disconnected:", rp.Identity())
}

func onActiveSpeakersChanged(ps []lksdk.Participant) {
	log.Println("Active speakers changed:")
}

func onRoomMetadataChanged(metaData string) {
	log.Println("Room metadata changed:", metaData)
}

func onReconnecting() {
	log.Println("Reconnecting:")
}

func onReconnected() {
	log.Println("Room reconnected:")
}

func onLocalTrackPublished(ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println(lp.Identity(), "Local track published:", ltp.Track().ID())
}

func onLocalTrackUnpublished(ltp *lksdk.LocalTrackPublication, lp *lksdk.LocalParticipant) {
	log.Println(lp.Identity(), "Local track unpublished:", ltp.Track().ID())
}

func onTrackMuted(tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println(p.Identity(), "Track muted:", tp.Track().ID())
}

func onTrackUnmuted(tp lksdk.TrackPublication, p lksdk.Participant) {
	log.Println(p.Identity(), "Track unmuted:", tp.Track().ID())
}

func onMetadataChanged(oldMetaData string, p lksdk.Participant) {
	log.Println(p.Identity(), "Metadata changed:", oldMetaData)
}

func onAttributesChanged(changed map[string]string, p lksdk.Participant) {
	log.Println(p.Identity(), "Attributes changed:", changed)
}

func onIsSpeakingChanged(p lksdk.Participant) {
	log.Println(p.Identity(), "Is speaking changed:")
}

func onConnectionQualityChanged(update *livekit.ConnectionQualityInfo, p lksdk.Participant) {
	log.Println(p.Identity(), "Connection quality changed:", update.String())
}

func onTrackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
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

func onTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unsubscribed:", track.ID(), publication.Track().ID(), rp.Identity())
}

func onTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscription failed:", sid, rp.Identity())
}

func onTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("OnTrackPublished", publication.MimeType(), rp.Identity())
}

func onTrackUnpublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unpublished:", publication.Track().ID(), rp.Identity())
}

func onDataPacket(data lksdk.DataPacket, params lksdk.DataReceiveParams) {
	userDataPackage, ok := data.(*lksdk.UserDataPacket)
	if ok {
		log.Println("Data packet received:", params.SenderIdentity, string(userDataPackage.Payload))
	}
}

func onTranscriptionReceived(transcriptionSegments []*lksdk.TranscriptionSegment, p lksdk.Participant, publication lksdk.TrackPublication) {
	log.Println("Transcription received:", transcriptionSegments, p, publication)
}

func NewRoomCallBack() *lksdk.RoomCallback {
	return &lksdk.RoomCallback{
		OnDisconnected:            onDisconnected,
		OnDisconnectedWithReason:  onDisconnectedWithReason,
		OnParticipantConnected:    onParticipantConnected,
		OnParticipantDisconnected: onParticipantDisconnected,
		OnActiveSpeakersChanged:   onActiveSpeakersChanged,
		OnRoomMetadataChanged:     onRoomMetadataChanged,
		OnReconnecting:            onReconnecting,
		OnReconnected:             onReconnected,
		ParticipantCallback: lksdk.ParticipantCallback{
			OnLocalTrackPublished:   onLocalTrackPublished,
			OnLocalTrackUnpublished: onLocalTrackUnpublished,
			// for all participants
			OnTrackMuted:               onTrackMuted,
			OnTrackUnmuted:             onTrackUnmuted,
			OnMetadataChanged:          onMetadataChanged,
			OnAttributesChanged:        onAttributesChanged,
			OnIsSpeakingChanged:        onIsSpeakingChanged,
			OnConnectionQualityChanged: onConnectionQualityChanged,

			// for remote participants
			OnTrackSubscribed:         onTrackSubscribed,
			OnTrackUnsubscribed:       onTrackUnsubscribed,
			OnTrackSubscriptionFailed: onTrackSubscriptionFailed,
			OnTrackPublished:          onTrackPublished,
			OnTrackUnpublished:        onTrackUnpublished,
			OnDataPacket:              onDataPacket,
			OnTranscriptionReceived:   onTranscriptionReceived,
		},
	}
}
