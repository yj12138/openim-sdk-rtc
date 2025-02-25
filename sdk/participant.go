package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"log"
)

type RemoteParticipant struct {
	lkRemoteParticipant *lksdk.RemoteParticipant
}

func NewRemoteParticipant(remoteParticipant *lksdk.RemoteParticipant) *RemoteParticipant {
	return &RemoteParticipant{
		lkRemoteParticipant: remoteParticipant,
	}
}

type LocalParticipant struct {
	lkLocalParticipant *lksdk.LocalParticipant
}

func NewLocalParticipant(localParticipant *lksdk.LocalParticipant) *LocalParticipant {
	return &LocalParticipant{
		lkLocalParticipant: localParticipant,
	}
}

func (p *LocalParticipant) PublicTrack(track *Track) {
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

func (p *LocalParticipant) UnpublishTrack(track *Track) {
	if track.liveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	err := p.lkLocalParticipant.UnpublishTrack(track.liveKitTrack.ID())
	if err != nil {
		log.Println(err)
	}
}

func (p *LocalParticipant) PublishData(topic string, data []byte, reliable bool, identifies []string) {
	p.lkLocalParticipant.PublishDataPacket(
		lksdk.UserData([]byte(data)),
		lksdk.WithDataPublishReliable(reliable),
		lksdk.WithDataPublishTopic(topic),
		lksdk.WithDataPublishDestination(identifies),
	)
}

func (p *LocalParticipant) SetSubscribed() {
}

func (p *LocalParticipant) SetMetadata(metaData string) {
	p.lkLocalParticipant.SetMetadata(metaData)
}

func (p *LocalParticipant) SetName(name string) {
	p.lkLocalParticipant.SetName(name)
}

func (p *LocalParticipant) SetAttributes(attributes map[string]string) {
	p.lkLocalParticipant.SetAttributes(attributes)
}

func (p *LocalParticipant) GetSessionStats() {
}
