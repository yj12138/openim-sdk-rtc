package sdk

import (
	"github.com/livekit/protocol/livekit"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type TranscriptionDataPacket struct {
	ParticipantIdentify string
	TrackId             string
	Segments            []*livekit.TranscriptionSegment
}

func (p *TranscriptionDataPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_Transcription{
		Transcription: &livekit.Transcription{
			TranscribedParticipantIdentity: p.ParticipantIdentify,
			TrackId:                        p.TrackId,
			Segments:                       p.Segments,
		},
	}}
}

func ConvertTranscriptionDataPacket(identify string, trackid string, _segments []*pb_room.TranscriptionSegment) *TranscriptionDataPacket {
	segments := make([]*livekit.TranscriptionSegment, len(_segments))
	for i, s := range _segments {
		segments[i] = &livekit.TranscriptionSegment{
			Id:        s.Id,
			Text:      s.Text,
			StartTime: s.StartTime,
			EndTime:   s.EndTime,
			Final:     s.Final,
			Language:  s.Language,
		}
	}
	return &TranscriptionDataPacket{
		ParticipantIdentify: identify,
		TrackId:             trackid,
		Segments:            segments,
	}
}

type ChatMessagePacket struct {
	Id            string
	Timestamp     int64
	EditTimestamp *int64
	Message       string
	Deleted       bool
	Generated     bool
}

func (p *ChatMessagePacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_ChatMessage{
		ChatMessage: &livekit.ChatMessage{},
	}}
}

type StreamHeaderPacket struct {
	StreamId       string
	Timestamp      int64
	Topic          string
	MimeType       string
	TotalLength    *uint64
	EncryptionType livekit.Encryption_Type
	Attributes     map[string]string
}

func (p *StreamHeaderPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_StreamHeader{
		StreamHeader: &livekit.DataStream_Header{
			StreamId:       p.StreamId,
			Timestamp:      p.Timestamp,
			Topic:          p.Topic,
			MimeType:       p.MimeType,
			TotalLength:    p.TotalLength,
			EncryptionType: p.EncryptionType,
			Attributes:     p.Attributes,
		},
	}}
}
func ConvertStreamHeaderPacket(header *pb_room.DataStream_Header) *StreamHeaderPacket {
	return &StreamHeaderPacket{
		StreamId:       header.StreamId,
		Timestamp:      header.Timestamp,
		Topic:          header.Topic,
		MimeType:       header.MimeType,
		TotalLength:    &header.TotalLength,
		EncryptionType: livekit.Encryption_NONE,
		Attributes:     header.Attributes,
	}
}

type StreamChunkPacket struct {
	StreamId   string
	ChunkIndex uint64
	Content    []byte
	Version    int32
	Iv         []byte
}

func (p *StreamChunkPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_StreamChunk{
		StreamChunk: &livekit.DataStream_Chunk{
			StreamId:   p.StreamId,
			ChunkIndex: p.ChunkIndex,
			Content:    p.Content,
			Version:    p.Version,
			Iv:         p.Iv,
		},
	}}
}
func ConvertStreamChunkPacket(chunk *pb_room.DataStream_Chunk) *StreamChunkPacket {
	return &StreamChunkPacket{
		StreamId:   chunk.StreamId,
		ChunkIndex: chunk.ChunkIndex,
		Content:    chunk.Content,
		Version:    chunk.Version,
		Iv:         chunk.Iv,
	}
}

type StreamTrailerPacket struct {
	StreamId   string
	Reason     string
	Attributes map[string]string
}

func (p *StreamTrailerPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_StreamTrailer{
		StreamTrailer: &livekit.DataStream_Trailer{
			StreamId:   p.StreamId,
			Reason:     p.Reason,
			Attributes: p.Attributes,
		},
	}}
}

func ConvertStreamTrailerPacket(trailer *pb_room.DataStream_Trailer) *StreamTrailerPacket {
	return &StreamTrailerPacket{
		StreamId:   trailer.StreamId,
		Reason:     trailer.Reason,
		Attributes: trailer.Attributes,
	}
}
