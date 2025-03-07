package sdk

import (
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"time"
)

type LocalParticipant struct {
	*lksdk.LocalParticipant
}

func NewLocalParticipant(localParticipant *lksdk.LocalParticipant) *LocalParticipant {
	return &LocalParticipant{
		LocalParticipant: localParticipant,
	}
}

func (p *LocalParticipant) SetSubscriptionPermissionWrap(allParticipants bool, permissions []*pb_track.ParticipantTrackPermission) {
	trackPermissions := make([]*livekit.TrackPermission, 0)
	for _, permission := range permissions {
		trackPermissions = append(trackPermissions, &livekit.TrackPermission{
			ParticipantIdentity: permission.ParticipantIdentity,
			AllTracks:           permission.AllowAll,
			TrackSids:           permission.AllowedTrackSids,
		})
	}
	sp := &livekit.SubscriptionPermission{
		AllParticipants:  allParticipants,
		TrackPermissions: trackPermissions,
	}
	p.SetSubscriptionPermission(sp)
}

func (p *LocalParticipant) SendData(topic string, data []byte, reliable bool, identifies []string) {
	p.PublishDataPacket(
		lksdk.UserData([]byte(data)),
		lksdk.WithDataPublishReliable(reliable),
		lksdk.WithDataPublishTopic(topic),
		lksdk.WithDataPublishDestination(identifies),
	)
}

func (p *LocalParticipant) SendTranscription(identify string, trackId string, segments []*pb_room.TranscriptionSegment) {
	packet := ConvertTranscriptionDataPacket(identify, trackId, segments)
	p.PublishDataPacket(packet)
}

func (p *LocalParticipant) SendStreamHeader(senderIdentity string, destinationIdentities []string, header *pb_room.DataStream_Header) {
	packet := ConvertStreamHeaderPacket(header)
	p.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendStreamChunk(senderIdentity string, destinationIdentities []string, chunk *pb_room.DataStream_Chunk) {
	packet := ConvertStreamChunkPacket(chunk)
	p.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendStreamTrailer(senderIdentity string, destinationIdentities []string, trailer *pb_room.DataStream_Trailer) {
	packet := ConvertStreamTrailerPacket(trailer)
	p.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendChatMessage(senderIdentity string, destinationIdentities []string, message string) {
	packet := &ChatMessagePacket{
		Id:        "",
		Timestamp: time.Now().Unix(),
		Message:   message,
		Deleted:   false,
		Generated: true,
	}
	p.PublishDataPacket(packet)
}
func (p *LocalParticipant) EditChatMessage(senderIdentity string, destinationIdentities []string, editText string, originalMessage *pb_room.ChatMessage) {
	packet := &ChatMessagePacket{
		Id:        "",
		Timestamp: time.Now().Unix(),
		Message:   editText,
		Deleted:   false,
		Generated: false,
	}
	p.PublishDataPacket(packet)
}
