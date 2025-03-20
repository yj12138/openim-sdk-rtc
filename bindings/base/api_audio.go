package base

import (
	"log"

	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	"github.com/openimsdk/openim-rtc/sdk"
	"github.com/openimsdk/openim-rtc/sdk/audio"
)

// Audio
func (api *API) NewAudioStream(req *pb_audio.NewAudioStreamRequest) *pb_audio.NewAudioStreamResponse {
	track := api.getRemoteTrack(req.TrackHandle)
	audioStream := sdk.NewAudioStreamByTrack(track, req.Type, req.SampleRate, req.NumChannels)
	streamHandle := api.storeObj(audioStream)
	audioStream.CallBack = func(audioFrame *sdk.AudioFrame) {
		buffer := NewAudioFrameBuffer(audioFrame.Payload, audioFrame.NumChannels, audioFrame.SampleRate, audioFrame.SamplesPerChannel)
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
	return res
}
func (api *API) NewAudioSource(req *pb_audio.NewAudioSourceRequest) *pb_audio.NewAudioSourceResponse {
	audioSource := sdk.NewAudioSource(req.Type, req.SampleRate, req.NumChannels)
	res := &pb_audio.NewAudioSourceResponse{
		Source: &pb_audio.OwnedAudioSource{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(audioSource)},
			Info: &pb_audio.AudioSourceInfo{
				Type: audioSource.SourceType,
			},
		},
	}
	return res
}
func (api *API) AudioFrameEchoCancellation(req *pb_audio.AudioFrameEchoCancellationRequest) *pb_audio.AudioFrameEchoCancellationResponse {
	rawLength := uint64(req.Buffer.NumChannels * req.Buffer.SamplesPerChannel * 2)
	echoLength := uint64(req.EchoBuffer.NumChannels * req.EchoBuffer.SamplesPerChannel * 2)
	if rawLength != echoLength || req.Buffer.SampleRate != req.EchoBuffer.SampleRate {
		log.Println("Audio Frame EchoCancle Error", req.Buffer, req.EchoBuffer)
		return nil
	}

	rawData := api.c.CPointerToGoByteSliceNoCopy(req.Buffer.DataPtr, rawLength)
	echoData := api.c.CPointerToGoByteSliceNoCopy(req.EchoBuffer.DataPtr, echoLength)
	sampleRate := req.Buffer.SampleRate
	endData := audio.GetAudioDSP().EchoCancellation(rawData, sampleRate, echoData)
	audioFrameBuffer := &AudioFrameBuffer{
		DataPtr:           api.c.GoByteSliceToCPointerNoCopy(endData),
		NumChannels:       req.Buffer.NumChannels,
		SampleRate:        req.Buffer.SampleRate,
		SamplesPerChannel: req.Buffer.SamplesPerChannel,
	}

	res := &pb_audio.AudioFrameEchoCancellationResponse{
		Buffer: &pb_audio.AudioFrameBufferInfo{
			DataPtr:           audioFrameBuffer.DataPtr,
			NumChannels:       audioFrameBuffer.NumChannels,
			SampleRate:        audioFrameBuffer.SampleRate,
			SamplesPerChannel: audioFrameBuffer.SamplesPerChannel,
		},
	}
	return res
}
func (api *API) CaptureAudioFrame(req *pb_audio.CaptureAudioFrameRequest) *pb_audio.CaptureAudioFrameResponse {
	asyncId := api.nextAsyncId()
	audioSource := api.getAudioSource(req.SourceHandle)
	go func() {
		buffer := req.Buffer
		sampleCount := buffer.NumChannels * buffer.SamplesPerChannel
		size := sampleCount * 2
		data := api.c.CPointerToGoByteSliceNoCopy(req.Buffer.DataPtr, uint64(size))
		audioSource.CaptureFrame(data, buffer.NumChannels, buffer.SampleRate, buffer.SamplesPerChannel)
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_CaptureAudioFrame{
				CaptureAudioFrame: &pb_audio.CaptureAudioFrameCallback{
					AsyncId: asyncId,
					Error:   "",
				},
			},
		})
	}()
	res := &pb_audio.CaptureAudioFrameResponse{AsyncId: asyncId}
	return res
}
func (api *API) ClearAudioBuffer(req *pb_audio.ClearAudioBufferRequest) *pb_audio.ClearAudioBufferResponse {
	audioSource := api.getAudioSource(req.SourceHandle)
	audioSource.ClearBuffer()
	res := &pb_audio.ClearAudioBufferResponse{}
	return res
}

func (api *API) NewAudioResampler(req *pb_audio.NewAudioResamplerRequest) *pb_audio.NewAudioResamplerResponse {
	resampler := sdk.NewAudioResampler()
	return &pb_audio.NewAudioResamplerResponse{
		Resampler: &pb_audio.OwnedAudioResampler{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(resampler)},
			Info:   &pb_audio.AudioResamplerInfo{},
		},
	}
}

func (api *API) RemixAndResample(req *pb_audio.RemixAndResampleRequest) *pb_audio.RemixAndResampleResponse {
	resample := api.GetAudioResampler(req.ResamplerHandle)
	length := uint64(req.Buffer.NumChannels * req.Buffer.SamplesPerChannel * 2)
	data := api.c.CPointerToGoByteSliceNoCopy(req.Buffer.DataPtr, length)
	buffer := resample.RemixAndResample(data, req.Buffer.SampleRate, req.Buffer.NumChannels, req.SampleRate, req.NumChannels)
	audioFrameBuffer := &AudioFrameBuffer{
		DataPtr:           api.c.GoByteSliceToCPointerNoCopy(buffer),
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
	}
}

func (api *API) AudioStreamFromParticipant(req *pb_audio.AudioStreamFromParticipantRequest) *pb_audio.AudioStreamFromParticipantResponse {
	return nil
}

func (api *API) NewSoxResampler(req *pb_audio.NewSoxResamplerRequest) *pb_audio.NewSoxResamplerResponse {
	return nil
}

func (api *API) PushSoxResampler(req *pb_audio.PushSoxResamplerRequest) *pb_audio.PushSoxResamplerResponse {
	return nil
}

func (api *API) FlushSoxResampler(req *pb_audio.FlushSoxResamplerRequest) *pb_audio.FlushSoxResamplerResponse {
	return nil
}

func (api *API) LoadAudioFilterPlugin(req *pb_audio.LoadAudioFilterPluginRequest) *pb_audio.LoadAudioFilterPluginResponse {
	return nil
}
