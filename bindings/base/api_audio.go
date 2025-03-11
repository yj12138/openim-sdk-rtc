package base

import (
	// lksdk "github.com/livekit/server-sdk-go/v2"
	"log"

	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	"github.com/openimsdk/openim-rtc/sdk"
)

type AudioFrameBuffer struct {
	DataPtr           uint64
	NumChannels       uint32
	SampleRate        uint32
	SamplesPerChannel uint32
}

// Audio
func (api *API) NewAudioStream(req *pb_audio.NewAudioStreamRequest) (*pb_audio.NewAudioStreamResponse, error) {
	track := api.getRemoteTrack(req.TrackHandle)
	audioStream := sdk.NewAudioStreamByTrack(track, req.Type, req.SampleRate, req.NumChannels)
	streamHandle := api.storeObj(audioStream)
	audioStream.CallBack = func(audioFrame *sdk.AudioFrame) {
		buffer := &AudioFrameBuffer{
			DataPtr:           goByteSliceToCPointerNoCopyFunc(audioFrame.Payload),
			NumChannels:       audioFrame.NumChannels,
			SampleRate:        audioFrame.SampleRate,
			SamplesPerChannel: audioFrame.SamplesPerChannel,
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_AudioStreamEvent{
				AudioStreamEvent: &pb_audio.AudioStreamEvent{
					StreamHandle: streamHandle,
					Message: &pb_audio.AudioStreamEvent_FrameReceived{
						FrameReceived: &pb_audio.AudioFrameReceived{
							Frame: &pb_audio.OwnedAudioFrameBuffer{
								Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(buffer)},
								Info: &pb_audio.AudioFrameBufferInfo{
									DataPtr:           buffer.DataPtr,
									SampleRate:        buffer.SampleRate,
									SamplesPerChannel: buffer.SamplesPerChannel,
									NumChannels:       buffer.NumChannels,
								},
							},
						},
					},
				},
			},
		})
	}
	res := &pb_audio.NewAudioStreamResponse{
		Stream: &pb_audio.OwnedAudioStream{
			Handle: &pb_handle.FfiOwnedHandle{Id: streamHandle},
			Info: &pb_audio.AudioStreamInfo{
				Type: audioStream.StreamType,
			},
		},
	}
	return res, nil
}
func (api *API) NewAudioSource(req *pb_audio.NewAudioSourceRequest) (*pb_audio.NewAudioSourceResponse, error) {
	audioSource := sdk.NewAudioSource(req.Type, req.SampleRate, req.NumChannels)
	res := &pb_audio.NewAudioSourceResponse{
		Source: &pb_audio.OwnedAudioSource{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(audioSource)},
			Info: &pb_audio.AudioSourceInfo{
				Type: audioSource.SourceType,
			},
		},
	}
	return res, nil
}
func (api *API) CaptureAudioFrame(req *pb_audio.CaptureAudioFrameRequest) (*pb_audio.CaptureAudioFrameResponse, error) {
	asyncId := api.nextAsyncId()
	audioSource := api.getAudioSource(req.SourceHandle)
	go func() {
		buffer := req.Buffer
		length := buffer.NumChannels * buffer.SampleRate * 2
		data := cPointerToGoByteSliceNoCopyFunc(req.Buffer.DataPtr, uint64(length))
		err := audioSource.CaptureFrame(data)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_CaptureAudioFrame{
				CaptureAudioFrame: &pb_audio.CaptureAudioFrameCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_audio.CaptureAudioFrameResponse{AsyncId: asyncId}
	return res, nil
}
func (api *API) ClearAudioBuffer(req *pb_audio.ClearAudioBufferRequest) (*pb_audio.ClearAudioBufferResponse, error) {
	audioSource := api.getAudioSource(req.SourceHandle)
	audioSource.ClearBuffer()
	res := &pb_audio.ClearAudioBufferResponse{}
	return res, nil
}

func (api *API) NewAudioResampler(req *pb_audio.NewAudioResamplerRequest) (*pb_audio.NewAudioResamplerResponse, error) {
	resampler := sdk.NewAudioResampler()
	return &pb_audio.NewAudioResamplerResponse{
		Resampler: &pb_audio.OwnedAudioResampler{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(resampler)},
			Info:   &pb_audio.AudioResamplerInfo{},
		},
	}, nil
}

func (api *API) RemixAndResample(req *pb_audio.RemixAndResampleRequest) (*pb_audio.RemixAndResampleResponse, error) {
	resample := api.GetAudioResampler(req.ResamplerHandle)
	length := uint64(req.Buffer.NumChannels * req.Buffer.SamplesPerChannel * 2)
	data := cPointerToGoByteSliceNoCopyFunc(req.Buffer.DataPtr, length)
	buffer, err := resample.RemixAndResample(data, req.Buffer.SampleRate, int(req.Buffer.NumChannels), req.SampleRate)
	if err != nil {
		log.Println("RemixAndResample: ", err.Error())
		return nil, err
	}
	audioFrameBuffer := &AudioFrameBuffer{
		DataPtr:           goByteSliceToCPointerNoCopyFunc(buffer),
		NumChannels:       req.NumChannels,
		SampleRate:        req.SampleRate,
		SamplesPerChannel: req.Buffer.SamplesPerChannel,
	}
	return &pb_audio.RemixAndResampleResponse{
		Buffer: &pb_audio.OwnedAudioFrameBuffer{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(audioFrameBuffer)},
			Info: &pb_audio.AudioFrameBufferInfo{
				DataPtr:           audioFrameBuffer.DataPtr,
				NumChannels:       audioFrameBuffer.NumChannels,
				SampleRate:        audioFrameBuffer.SampleRate,
				SamplesPerChannel: audioFrameBuffer.SamplesPerChannel,
			},
		},
	}, nil
}

func (api *API) AudioStreamFromParticipant(req *pb_audio.AudioStreamFromParticipantRequest) (*pb_audio.AudioStreamFromParticipantResponse, error) {
	return nil, nil
}

func (api *API) NewSoxResampler(req *pb_audio.NewSoxResamplerRequest) (*pb_audio.NewSoxResamplerResponse, error) {
	return nil, nil
}

func (api *API) PushSoxResampler(req *pb_audio.PushSoxResamplerRequest) (*pb_audio.PushSoxResamplerResponse, error) {
	return nil, nil
}

func (api *API) FlushSoxResampler(req *pb_audio.FlushSoxResamplerRequest) (*pb_audio.FlushSoxResamplerResponse, error) {
	return nil, nil
}

func (api *API) LoadAudioFilterPlugin(req *pb_audio.LoadAudioFilterPluginRequest) (*pb_audio.LoadAudioFilterPluginResponse, error) {
	return nil, nil
}
