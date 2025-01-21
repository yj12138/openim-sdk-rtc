package ui

import (
	"log"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
)

type RoomCallBack struct {
	winRoom *WindowRoom
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
		// for {
		// pkt, _, err := track.ReadRTP()
		// if err != nil {
		// 	break
		// }
		// }
	}()
}

func OnTrackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println("Track unsubscribed:", track.ID(), publication.Track().ID(), rp.Identity())
}

func OnTrackSubscriptionFailed(sid string, rp *lksdk.RemoteParticipant) {
	log.Println("Track subscription failed:", sid, rp.Identity())
}

func OnTrackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	log.Println(publication, rp)
	// if publication == nil {
	// 	log.Println("publication  is nil")
	// 	return
	// }
	// if rp == nil {
	// 	log.Println("rp is nil")
	// 	return
	// }
	// log.Println("Track published:", publication.Track().ID(), rp.Identity())

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

func NewRoomCallBack(winRoom *WindowRoom) *lksdk.RoomCallback {
	callBack := &RoomCallBack{
		winRoom: winRoom,
	}
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
