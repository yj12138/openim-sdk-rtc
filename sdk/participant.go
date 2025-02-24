package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"log"
)

type Participant struct {
	lkLocalParticipant *lksdk.LocalParticipant
}

func NewParticipant(localParticipant *lksdk.LocalParticipant) *Participant {
	return &Participant{
		lkLocalParticipant: localParticipant,
	}
}

func (p *Participant) PublicTrack(track *Track) {
	if track.liveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	trackPublication, err := p.lkLocalParticipant.PublishTrack(track.liveKitTrack, &lksdk.TrackPublicationOptions{
		VideoWidth:  track.videoWidth,
		VideoHeight: track.videoHeight,
		Name:        track.name,
	})
	track.localTrackPublication = trackPublication
	if err != nil {
		log.Panic(err)
	}
}

func (p *Participant) UnpublishTrack(track *Track) {
	if track.liveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	err := p.lkLocalParticipant.UnpublishTrack(track.liveKitTrack.ID())
	if err != nil {
		log.Println(err)
	}
}

func (p *Participant) PublishData(topic string, data string, reliable bool, identifies []string) {
	p.lkLocalParticipant.PublishDataPacket(
		lksdk.UserData([]byte(data)),
		lksdk.WithDataPublishReliable(reliable),
		lksdk.WithDataPublishTopic(topic),
		lksdk.WithDataPublishDestination(identifies),
	)
}

func (p *Participant) SetSubscribed() {
}

func (p *Participant) UpdateLocalMetadata() {
}

func (p *Participant) UpdateLocalName() {
}

func (p *Participant) GetSessionStats() {
}
