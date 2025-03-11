package sdk

import (
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	"github.com/pion/webrtc/v4"
)

type AudioFrame struct {
	Payload           []byte
	SampleRate        uint32
	NumChannels       uint32
	SamplesPerChannel uint32
}

type AudioStream struct {
	track       *webrtc.TrackRemote
	StreamType  pb_audio.AudioStreamType
	SampleRate  uint32
	NumChannels uint32
	CallBack    func(frame *AudioFrame)
}

func (s *AudioStream) handleStream() {
	for {
		pkt, _, err := s.track.ReadRTP()
		if err != nil {
			break
		}
		if s.CallBack != nil {
			s.CallBack(&AudioFrame{
				Payload:           pkt.Payload,
				SampleRate:        s.SampleRate,
				NumChannels:       s.NumChannels,
				SamplesPerChannel: uint32(len(pkt.Payload) / int(s.NumChannels)),
			})
		}
	}
}

func NewAudioStreamByTrack(track *webrtc.TrackRemote, streamType pb_audio.AudioStreamType, sampleRate uint32, numChannels uint32) *AudioStream {
	stream := &AudioStream{
		track:       track,
		StreamType:  streamType,
		SampleRate:  sampleRate,
		NumChannels: numChannels,
	}
	go stream.handleStream()
	return stream
}
