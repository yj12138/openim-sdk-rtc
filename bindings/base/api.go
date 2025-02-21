package base

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/proto/go/audio"
	"github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/openimsdk/openim-rtc/proto/go/video"
	"github.com/openimsdk/openim-rtc/sdk"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

// room
func (api *API) Connect(req *room.ConnectReq) (*room.ConnectRes, error) {
	sdk.NewRoom(lksdk.NewRoomCallback())
	return nil, nil
}
func (api *API) Disconnect(req *room.DisconnectReq) (*room.DisconnectRes, error) {

	return nil, nil
}
func (api *API) PublishTrack(req *room.PublishTrackReq) (*room.PublishTrackRes, error) {

	return nil, nil
}
func (api *API) UnpublishTrack(req *room.UnpublishTrackReq) (*room.UnpublishTrackRes, error) {

	return nil, nil
}
func (api *API) PublishData(req *room.PublishDataReq) (*room.PublishDataRes, error) {

	return nil, nil
}
func (api *API) SetSubscribed(req *room.SetSubscribedReq) (*room.SetSubscribedRes, error) {

	return nil, nil
}
func (api *API) SetLocalMetadata(req *room.SetLocalMetadataReq) (*room.SetLocalMetadataRes, error) {

	return nil, nil
}
func (api *API) SetLocalName(req *room.SetLocalNameReq) (*room.SetLocalNameRes, error) {

	return nil, nil
}
func (api *API) SetLocalAttributes(req *room.SetLocalAttributesReq) (*room.SetLocalAttributesRes, error) {

	return nil, nil
}
func (api *API) GetSessionStats(req *room.GetSessionStatsReq) (*room.GetSessionStatsRes, error) {

	return nil, nil
}
func (api *API) PublishTranscription(req *room.PublishTranscriptionReq) (*room.PublishTranscriptionRes, error) {

	return nil, nil
}
func (api *API) EnableRemoteTrackPublication(req *room.EnableRemoteTrackPublicationReq) (*room.EnableRemoteTrackPublicationRes, error) {

	return nil, nil
}
func (api *API) UpdateRemoteTrackPublicationDimension(req *room.UpdateRemoteTrackPublicationDimensionReq) (*room.UpdateRemoteTrackPublicationDimensionRes, error) {

	return nil, nil
}
func (api *API) SendStreamHeader(req *room.SendStreamHeaderReq) (*room.SendStreamHeaderRes, error) {

	return nil, nil
}
func (api *API) SendStreamChunk(req *room.SendStreamChunkReq) (*room.SendStreamChunkRes, error) {

	return nil, nil
}
func (api *API) SendStreamTrailer(req *room.SendStreamTrailerReq) (*room.SendStreamTrailerRes, error) {

	return nil, nil
}
func (api *API) SetDataChannelBufferedAmountLowThreshold(req *room.SetDataChannelBufferedAmountLowThresholdReq) (*room.SetDataChannelBufferedAmountLowThresholdRes, error) {

	return nil, nil
}

// Track
func (api *API) CreateVideoTrack(req *track.CreateVideoTrackReq) (*track.CreateAudioTrackRes, error) {

	return nil, nil
}
func (api *API) CreateAudioTrack(req *track.CreateAudioTrackReq) (*track.CreateAudioTrackRes, error) {

	return nil, nil
}
func (api *API) LocalTrackMute(req *track.LocalTrackMuteReq) (*track.LocalTrackMuteRes, error) {

	return nil, nil
}
func (api *API) EnableRemoteTrack(req *track.EnableRemoteTrackReq) (*track.EnableRemoteTrackRes, error) {

	return nil, nil
}
func (api *API) GetStats(req *track.GetStatsReq) (*track.GetStatsRes, error) {

	return nil, nil
}
func (api *API) SetTrackSubscriptionPermissions(req *track.SetTrackSubscriptionPermissionsReq) (*track.SetTrackSubscriptionPermissionsRes, error) {

	return nil, nil
}

// Audio
func (api *API) NewAudioStream(req *audio.NewAudioStreamReq) (*audio.NewAudioStreamRes, error) {
	return nil, nil
}
func (api *API) NewAudioSource(req *audio.NewAudioSourceReq) (*audio.NewAudioSourceRes, error) {

	return nil, nil
}
func (api *API) CaptureAudioFrame(req *audio.CaptureAudioFrameReq) (*audio.CaptureAudioFrameRes, error) {

	return nil, nil
}
func (api *API) ClearAudioBuffer(req *audio.ClearAudioBufferReq) (*audio.ClearAudioBufferRes, error) {

	return nil, nil
}
func (api *API) NewAudioResampler(req *audio.NewAudioResamplerReq) (*audio.NewAudioResamplerRes, error) {

	return nil, nil
}
func (api *API) RemixAndResample(req *audio.RemixAndResampleReq) (*audio.RemixAndResampleRes, error) {

	return nil, nil
}
func (api *API) AudioStreamFromParticipant(req *audio.AudioStreamFromParticipantReq) (*audio.AudioStreamFromParticipantRes, error) {

	return nil, nil
}

// Video
func (api *API) NewVideoStream(req *video.NewVideoStreamReq) (*video.NewVideoStreamRes, error) {
	return nil, nil
}
func (api *API) NewVideoSource(req *video.NewVideoSourceReq) (*video.NewVideoSourceRes, error) {

	return nil, nil
}

func (api *API) CaptureVideoFrame(req *video.CaptureVideoFrameReq) (*video.CaptureVideoFrameRes, error) {

	return nil, nil
}

func (api *API) VideoConvert(req *video.VideoConvertReq) (*video.VideoConvertRes, error) {

	return nil, nil
}
func (api *API) VideoStreamFromParticipan(req *video.VideoStreamFromParticipanReq) (*video.VideoStreamFromParticipanRes, error) {

	return nil, nil
}
