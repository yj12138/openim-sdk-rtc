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

func (p *LocalParticipant) SendData(topic string, data []byte, reliable bool, identifies []string) error {
	packet := lksdk.UserData([]byte(data))
	return p.PublishDataPacket(
		packet,
		lksdk.WithDataPublishReliable(reliable),
		lksdk.WithDataPublishTopic(topic),
		lksdk.WithDataPublishDestination(identifies),
	)
}

func (p *LocalParticipant) SendSipDtmf(code uint32, digit string) error {
	packet := &SipDTMFPacket{
		Code:  code,
		Digit: digit,
	}
	return p.PublishDataPacket(packet)
}

func (p *LocalParticipant) SendTranscription(identify string, trackId string, segments []*pb_room.TranscriptionSegment) error {
	_segments := make([]*livekit.TranscriptionSegment, len(segments))
	for i, s := range _segments {
		_segments[i] = &livekit.TranscriptionSegment{
			Id:        s.Id,
			Text:      s.Text,
			StartTime: s.StartTime,
			EndTime:   s.EndTime,
			Final:     s.Final,
			Language:  s.Language,
		}
	}
	packet := &TranscriptionDataPacket{
		ParticipantIdentify: identify,
		TrackId:             trackId,
		Segments:            _segments,
	}
	return p.PublishDataPacket(packet)
}

func (p *LocalParticipant) SendStreamHeader(senderIdentity string, destinationIdentities []string, header *pb_room.DataStream_Header) error {
	packet := &StreamHeaderPacket{
		StreamId:       header.StreamId,
		Timestamp:      header.Timestamp,
		Topic:          header.Topic,
		MimeType:       header.MimeType,
		TotalLength:    &header.TotalLength,
		EncryptionType: livekit.Encryption_NONE,
		Attributes:     header.Attributes,
	}
	return p.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendStreamChunk(senderIdentity string, destinationIdentities []string, chunk *pb_room.DataStream_Chunk) error {
	packet := &StreamChunkPacket{
		StreamId:   chunk.StreamId,
		ChunkIndex: chunk.ChunkIndex,
		Content:    chunk.Content,
		Version:    chunk.Version,
		Iv:         chunk.Iv,
	}
	return p.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendStreamTrailer(senderIdentity string, destinationIdentities []string, trailer *pb_room.DataStream_Trailer) error {
	packet := &StreamTrailerPacket{
		StreamId:   trailer.StreamId,
		Reason:     trailer.Reason,
		Attributes: trailer.Attributes,
	}
	return p.PublishDataPacket(packet)
}
func (p *LocalParticipant) SendChatMessage(senderIdentity string, destinationIdentities []string, message string) error {
	packet := &ChatMessagePacket{
		Id:        "",
		Timestamp: time.Now().Unix(),
		Message:   message,
		Deleted:   false,
		Generated: true,
	}
	return p.PublishDataPacket(packet)
}
func (p *LocalParticipant) EditChatMessage(senderIdentity string, destinationIdentities []string, editText string, originalMessage *pb_room.ChatMessage) error {
	packet := &ChatMessagePacket{
		Id:        "",
		Timestamp: time.Now().Unix(),
		Message:   editText,
		Deleted:   false,
		Generated: false,
	}
	return p.PublishDataPacket(packet)
}

func (p *LocalParticipant) PerformRpc(identify string, method string, payload string, responseTimeoutMs uint32) {
	packet := &RpcRequestPacket{
		Id:                identify,
		Method:            method,
		Payload:           payload,
		ResponseTimeoutMs: responseTimeoutMs,
	}
	p.PublishDataPacket(packet)
}
