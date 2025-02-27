package base

import (
	// lksdk "github.com/livekit/server-sdk-go/v2"
	pb_audio "github.com/openimsdk/openim-rtc/proto/go/audio"
	"github.com/openimsdk/openim-rtc/sdk"
)

// Audio
func (api *API) NewAudioStream(req *pb_audio.NewAudioStreamReq) (*pb_audio.NewAudioStreamRes, error) {
	track := api.getTrack(req.TrackHandle)
	audioStream := sdk.NewAudioStreamByTrack(track, req.Type, req.SampleRate, req.NumChannels)
	streamHandle := api.storeObj(audioStream)
	res := &pb_audio.NewAudioStreamRes{
		Stream: &pb_audio.OwnedAudioStream{
			Handle: streamHandle,
			Info: &pb_audio.AudioStreamInfo{
				Type: audioStream.StreamType,
			},
		},
	}
	return res, nil
}
func (api *API) NewAudioSource(req *pb_audio.NewAudioSourceReq) (*pb_audio.NewAudioSourceRes, error) {

	return nil, nil
}
func (api *API) CaptureAudioFrame(req *pb_audio.CaptureAudioFrameReq) (*pb_audio.CaptureAudioFrameRes, error) {

	return nil, nil
}
func (api *API) ClearAudioBuffer(req *pb_audio.ClearAudioBufferReq) (*pb_audio.ClearAudioBufferRes, error) {

	return nil, nil
}
func (api *API) NewAudioResampler(req *pb_audio.NewAudioResamplerReq) (*pb_audio.NewAudioResamplerRes, error) {

	return nil, nil
}
func (api *API) RemixAndResample(req *pb_audio.RemixAndResampleReq) (*pb_audio.RemixAndResampleRes, error) {

	return nil, nil
}
func (api *API) AudioStreamFromParticipant(req *pb_audio.AudioStreamFromParticipantReq) (*pb_audio.AudioStreamFromParticipantRes, error) {

	return nil, nil
}
