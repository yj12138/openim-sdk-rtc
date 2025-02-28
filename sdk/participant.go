package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"log"
)

type LocalParticipant struct {
	LiveKitLocalParticipant *lksdk.LocalParticipant
}

func NewLocalParticipant(localParticipant *lksdk.LocalParticipant) *LocalParticipant {
	return &LocalParticipant{
		LiveKitLocalParticipant: localParticipant,
	}
}

func (p *LocalParticipant) PublicTrack(track *LocalTrack) {
	if track.LiveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	trackPublication, err := p.LiveKitLocalParticipant.PublishTrack(track.LiveKitTrack, &lksdk.TrackPublicationOptions{
		VideoWidth:  track.VideoWidth,
		VideoHeight: track.VideoHeight,
		Name:        track.Name,
	})
	track.LocalTrackPublication = trackPublication
	if err != nil {
		log.Panic(err)
	}
}

func (p *LocalParticipant) UnpublishTrack(track *LocalTrack) {
	if track.LiveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	err := p.LiveKitLocalParticipant.UnpublishTrack(track.LocalTrackPublication.SID())
	if err != nil {
		log.Println(err)
	}
}

func (p *LocalParticipant) PublishData(topic string, data []byte, reliable bool, identifies []string) {
	p.LiveKitLocalParticipant.PublishDataPacket(
		lksdk.UserData([]byte(data)),
		lksdk.WithDataPublishReliable(reliable),
		lksdk.WithDataPublishTopic(topic),
		lksdk.WithDataPublishDestination(identifies),
	)
}

func (p *LocalParticipant) SetSubscribed() {
}

func (p *LocalParticipant) SetMetadata(metaData string) {
	p.LiveKitLocalParticipant.SetMetadata(metaData)
}

func (p *LocalParticipant) SetName(name string) {
	p.LiveKitLocalParticipant.SetName(name)
}

func (p *LocalParticipant) SetAttributes(attributes map[string]string) {
	p.LiveKitLocalParticipant.SetAttributes(attributes)
}

type RemoteParticipant struct {
	LiveKitRemoteParticipant *lksdk.RemoteParticipant
}

func NewRemoteParticipant(remoteParticipant *lksdk.RemoteParticipant) *RemoteParticipant {
	return &RemoteParticipant{
		LiveKitRemoteParticipant: remoteParticipant,
	}
}
