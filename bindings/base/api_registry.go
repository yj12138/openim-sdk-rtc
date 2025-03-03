package base

import (
	"github.com/openimsdk/openim-rtc/proto/go/event"
)

var funcMap = map[event.FuncEventName]callFunc{
	// room
	event.FuncEventName_Connect:         wrapFunc(api.Connect),
	event.FuncEventName_Disconnect:      wrapFunc(api.Disconnect),
	event.FuncEventName_GetConnectState: wrapFunc(api.GetConnectState),
	// participant
	event.FuncEventName_PublishTrack:                          wrapFunc(api.PublishTrack),
	event.FuncEventName_UnpublishTrack:                        wrapFunc(api.UnpublishTrack),
	event.FuncEventName_PublishData:                           wrapFunc(api.PublishData),
	event.FuncEventName_SetSubscribed:                         wrapFunc(api.SetSubscribed),
	event.FuncEventName_SetLocalMetadata:                      wrapFunc(api.SetLocalMetadata),
	event.FuncEventName_SetLocalName:                          wrapFunc(api.SetLocalName),
	event.FuncEventName_SetLocalAttributes:                    wrapFunc(api.SetLocalAttributes),
	event.FuncEventName_PublishTranscription:                  wrapFunc(api.PublishTranscription),
	event.FuncEventName_EnableRemoteTrackPublication:          wrapFunc(api.EnableRemoteTrackPublication),
	event.FuncEventName_UpdateRemoteTrackPublicationDimension: wrapFunc(api.UpdateRemoteTrackPublicationDimension),
	event.FuncEventName_SendStreamHeader:                      wrapFunc(api.SendStreamHeader),
	event.FuncEventName_SendStreamChunk:                       wrapFunc(api.SendStreamChunk),
	event.FuncEventName_SendStreamTrailer:                     wrapFunc(api.SendStreamTrailer),
	event.FuncEventName_SendChatMessageRequest:                wrapFunc(api.SendChatMessage),
	event.FuncEventName_EditChatMessageRequest:                wrapFunc(api.EditChatMessage),
	// track
	event.FuncEventName_CreateVideoTrack:                wrapFunc(api.CreateVideoTrack),
	event.FuncEventName_CreateAudioTrack:                wrapFunc(api.CreateAudioTrack),
	event.FuncEventName_LocalTrackMute:                  wrapFunc(api.LocalTrackMute),
	event.FuncEventName_EnableRemoteTrack:               wrapFunc(api.EnableRemoteTrack),
	event.FuncEventName_SetTrackSubscriptionPermissions: wrapFunc(api.SetTrackSubscriptionPermissions),
	// audio
	event.FuncEventName_NewAudioStream:             wrapFunc(api.NewAudioStream),
	event.FuncEventName_NewAudioSource:             wrapFunc(api.NewAudioSource),
	event.FuncEventName_CaptureAudioFrame:          wrapFunc(api.CaptureAudioFrame),
	event.FuncEventName_ClearAudioBuffer:           wrapFunc(api.ClearAudioBuffer),
	event.FuncEventName_NewAudioResampler:          wrapFunc(api.NewAudioResampler),
	event.FuncEventName_RemixAndResample:           wrapFunc(api.RemixAndResample),
	event.FuncEventName_AudioStreamFromParticipant: wrapFunc(api.AudioStreamFromParticipant),
	// video
	event.FuncEventName_NewVideoStream:            wrapFunc(api.NewVideoStream),
	event.FuncEventName_NewVideoSource:            wrapFunc(api.NewVideoSource),
	event.FuncEventName_CaptureVideoFrame:         wrapFunc(api.CaptureVideoFrame),
	event.FuncEventName_VideoConvert:              wrapFunc(api.VideoConvert),
	event.FuncEventName_VideoStreamFromParticipan: wrapFunc(api.VideoStreamFromParticipan),
}
