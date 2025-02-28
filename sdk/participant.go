package sdk

import (
	"log"
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"

	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
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

func (p *LocalParticipant) SendData(topic string, data []byte, reliable bool, identifies []string) {
	p.LiveKitLocalParticipant.PublishDataPacket(
		lksdk.UserData([]byte(data)),
		lksdk.WithDataPublishReliable(reliable),
		lksdk.WithDataPublishTopic(topic),
		lksdk.WithDataPublishDestination(identifies),
	)
}

func (p *LocalParticipant) SendTranscription(identify string, trackId string, segments []*pb_participant.TranscriptionSegment) {
	packet := ConvertTranscriptionDataPacket(identify, trackId, segments)
	p.LiveKitLocalParticipant.PublishDataPacket(packet)
}

func (p *LocalParticipant) SendStreamHeader(senderIdentity string, destinationIdentities []string, header *pb_participant.DataStream_Header) {
	packet := ConvertStreamHeaderPacket(header)
	p.LiveKitLocalParticipant.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendStreamChunk(senderIdentity string, destinationIdentities string, chunk *pb_participant.DataStream_Chunk) {
	packet := ConvertStreamChunkPacket(chunk)
	p.LiveKitLocalParticipant.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendStreamTrailer(senderIdentity string, destinationIdentities string, trailer *pb_participant.DataStream_Trailer) {
	packet := ConvertStreamTrailerPacket(trailer)
	p.LiveKitLocalParticipant.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendChatMessage(senderIdentity string, destinationIdentities string, message string) {
	packet := &ChatMessagePacket{
		Id:        "",
		Timestamp: time.Now().Unix(),
		Message:   message,
		Deleted:   false,
		Generated: true,
	}
	p.LiveKitLocalParticipant.PublishDataPacket(packet)
}
func (p *LocalParticipant) EditChatMessage(senderIdentity string, destinationIdentities string, editText string, originalMessage *pb_participant.ChatMessage) {
	packet := &ChatMessagePacket{
		Id:        "",
		Timestamp: time.Now().Unix(),
		Message:   editText,
		Deleted:   false,
		Generated: false,
	}
	p.LiveKitLocalParticipant.PublishDataPacket(packet)
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
	remoteTrackPublications  map[string]*lksdk.RemoteTrackPublication
}

func NewRemoteParticipant(remoteParticipant *lksdk.RemoteParticipant) *RemoteParticipant {
	return &RemoteParticipant{
		LiveKitRemoteParticipant: remoteParticipant,
		remoteTrackPublications:  make(map[string]*lksdk.RemoteTrackPublication),
	}
}

func (rp *RemoteParticipant) addRemoteTrackPublication(rtp *lksdk.RemoteTrackPublication) {
	_, ok := rp.remoteTrackPublications[rtp.SID()]
	if !ok {
		rp.remoteTrackPublications[rtp.SID()] = rtp
	}
}

func (rp *RemoteParticipant) SetSubscribed(publicationId string, subscribe bool) {
	publication, ok := rp.remoteTrackPublications[publicationId]
	if ok {
		publication.SetSubscribed(subscribe)
	}
}
