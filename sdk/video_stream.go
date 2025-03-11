package sdk

import (
	pb_video "github.com/openimsdk/openim-rtc/proto/go/video_frame"
	"github.com/pion/webrtc/v4"
)

type VideoFrame struct {
	Payload     []byte
	Rotation    pb_video.VideoRotation
	TimeStampUs int64
	Type        pb_video.VideoBufferType
	Width       uint32
	Height      uint32
}

type VideoStream struct {
	track      *webrtc.TrackRemote
	StreamType pb_video.VideoStreamType
	CallBack   func(frame *VideoFrame)
}

func (s *VideoStream) handleStream() {
	for {
		pkt, _, err := s.track.ReadRTP()
		if err != nil {
			break
		}
		if s.CallBack != nil {
			s.CallBack(&VideoFrame{
				Payload: pkt.Payload,
			})
		}
	}
}

func NewVideoStreamByTrack(track *webrtc.TrackRemote, streamType pb_video.VideoStreamType, format pb_video.VideoBufferType, normalizeStride bool) *VideoStream {
	stream := &VideoStream{
		track:      track,
		StreamType: streamType,
	}
	go stream.handleStream()
	return stream
}
