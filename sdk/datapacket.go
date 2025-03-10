package sdk

import (
	"github.com/livekit/protocol/livekit"
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

type SipDTMFPacket struct {
	Code  uint32
	Digit string
}

func (s *SipDTMFPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_SipDtmf{
		SipDtmf: &livekit.SipDTMF{
			Code:  s.Code,
			Digit: s.Digit,
		},
	}}
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

type RpcRequestPacket struct {
	Id                string
	Method            string
	Payload           string
	ResponseTimeoutMs uint32
	Version           uint32
}

func (p *RpcRequestPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_RpcRequest{
		RpcRequest: &livekit.RpcRequest{
			Id:                p.Id,
			Method:            p.Method,
			Payload:           p.Payload,
			ResponseTimeoutMs: p.ResponseTimeoutMs,
			Version:           p.Version,
		},
	}}
}

type RpcAckPacket struct {
	RequestId string
}

func (p *RpcAckPacket) ToProto() *livekit.DataPacket {
	return &livekit.DataPacket{Value: &livekit.DataPacket_RpcAck{
		RpcAck: &livekit.RpcAck{
			RequestId: p.RequestId,
		},
	}}
}

type RpcResponsePacket struct {
	RequestId string
	Payload   string
	Code      uint32
	Message   string
	Data      string
}

func (p *RpcResponsePacket) ToProto() *livekit.DataPacket {
	if p.Code > 0 {
		return &livekit.DataPacket{Value: &livekit.DataPacket_RpcResponse{
			RpcResponse: &livekit.RpcResponse{
				RequestId: p.RequestId,
				Value: &livekit.RpcResponse_Payload{
					Payload: p.Payload,
				},
			},
		}}
	} else {
		return &livekit.DataPacket{Value: &livekit.DataPacket_RpcResponse{
			RpcResponse: &livekit.RpcResponse{
				RequestId: p.RequestId,
				Value: &livekit.RpcResponse_Error{
					Error: &livekit.RpcError{
						Code:    p.Code,
						Message: p.Message,
						Data:    p.Data,
					},
				},
			},
		}}
	}

}
