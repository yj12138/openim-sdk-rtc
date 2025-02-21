package base

import (
	"github.com/openimsdk/openim-rtc/proto/go/event"
	"google.golang.org/protobuf/proto"
	// "log"
	"time"
)

var (
	api *API
)

func init() {
	api = NewAPI()
}

func wrapFunc[A, B proto.Message](fn func(req A) (B, error)) callFunc {
	return func(handleId int64, name event.FuncRequestEventName, reqData []byte) (resData []byte, err error) {
		start := time.Now()
		var req A
		var res B

		if err := proto.Unmarshal(reqData, req); err != nil {
			return nil, err
		}

		defer func(start time.Time) {
			if r := recover(); r != nil {
				// err = sdkerrs.ErrPanic(r)
				// log.ZPanic(ctx, "wrapFunc recover", err)
			}
			// if _, ignored := ignoredLogFuncMap[name]; ignored {
			// 	return
			// }
			// elapsed := time.Since(start)
			// if err == nil {
			// 	log.ZInfo(ctx, fmt.Sprintf("[%s -> Go SDK] Response Success - %s", sdkcommon.AppFramework_name[int32(open_im_sdk.IMUserContext.Info().AppFramework)],
			// 		sdkevent.FuncRequestEventName_name[int32(name)]), "duration", elapsed, "resp", pbResp)
			// } else {
			// 	log.ZError(ctx, fmt.Sprintf("[%s -> Go SDK] Response Error - %s", sdkcommon.AppFramework_name[int32(open_im_sdk.IMUserContext.Info().AppFramework)],
			// 		sdkevent.FuncRequestEventName_name[int32(name)]), err, "duration", elapsed)

			// }
		}(start)

		// if _, ignored := ignoredLogFuncMap[name]; !ignored {
		// 	log.ZInfo(ctx, fmt.Sprintf("[%s -> Go SDK]  Request - %s", sdkcommon.AppFramework_name[int32(open_im_sdk.IMUserContext.Info().AppFramework)],
		// 		sdkevent.FuncRequestEventName_name[int32(name)]), "req", pbReq)
		// }

		// if err := checker.Validate(&pbReq); err != nil {
		// 	return nil, err
		// }

		res, err = fn(req)
		if err != nil {
			return nil, err
		}
		return proto.Marshal(res)
	}
}

var funcMap = map[event.FuncRequestEventName]callFunc{
	// room
	event.FuncRequestEventName_Connect:                                  wrapFunc(api.Connect),
	event.FuncRequestEventName_Disconnect:                               wrapFunc(api.Disconnect),
	event.FuncRequestEventName_PublishTrack:                             wrapFunc(api.PublishTrack),
	event.FuncRequestEventName_UnpublishTrack:                           wrapFunc(api.UnpublishTrack),
	event.FuncRequestEventName_PublishData:                              wrapFunc(api.PublishData),
	event.FuncRequestEventName_SetSubscribed:                            wrapFunc(api.SetSubscribed),
	event.FuncRequestEventName_SetLocalMetadata:                         wrapFunc(api.SetLocalMetadata),
	event.FuncRequestEventName_SetLocalName:                             wrapFunc(api.SetLocalName),
	event.FuncRequestEventName_SetLocalAttributes:                       wrapFunc(api.SetLocalAttributes),
	event.FuncRequestEventName_GetSessionStats:                          wrapFunc(api.GetSessionStats),
	event.FuncRequestEventName_PublishTranscription:                     wrapFunc(api.PublishTranscription),
	event.FuncRequestEventName_EnableRemoteTrackPublication:             wrapFunc(api.EnableRemoteTrackPublication),
	event.FuncRequestEventName_UpdateRemoteTrackPublicationDimension:    wrapFunc(api.UpdateRemoteTrackPublicationDimension),
	event.FuncRequestEventName_SendStreamHeader:                         wrapFunc(api.SendStreamHeader),
	event.FuncRequestEventName_SendStreamChunk:                          wrapFunc(api.SendStreamChunk),
	event.FuncRequestEventName_SendStreamTrailer:                        wrapFunc(api.SendStreamTrailer),
	event.FuncRequestEventName_SetDataChannelBufferedAmountLowThreshold: wrapFunc(api.SetDataChannelBufferedAmountLowThreshold),
	// track
	event.FuncRequestEventName_CreateVideoTrack:                wrapFunc(api.CreateVideoTrack),
	event.FuncRequestEventName_CreateAudioTrack:                wrapFunc(api.CreateAudioTrack),
	event.FuncRequestEventName_LocalTrackMute:                  wrapFunc(api.LocalTrackMute),
	event.FuncRequestEventName_EnableRemoteTrack:               wrapFunc(api.EnableRemoteTrack),
	event.FuncRequestEventName_GetStats:                        wrapFunc(api.GetStats),
	event.FuncRequestEventName_SetTrackSubscriptionPermissions: wrapFunc(api.SetTrackSubscriptionPermissions),
	// audio
	event.FuncRequestEventName_NewAudioStream:             wrapFunc(api.NewAudioStream),
	event.FuncRequestEventName_NewAudioSource:             wrapFunc(api.NewAudioSource),
	event.FuncRequestEventName_CaptureAudioFrame:          wrapFunc(api.CaptureAudioFrame),
	event.FuncRequestEventName_ClearAudioBuffer:           wrapFunc(api.ClearAudioBuffer),
	event.FuncRequestEventName_NewAudioResampler:          wrapFunc(api.NewAudioResampler),
	event.FuncRequestEventName_RemixAndResample:           wrapFunc(api.RemixAndResample),
	event.FuncRequestEventName_AudioStreamFromParticipant: wrapFunc(api.AudioStreamFromParticipant),
	// video
	event.FuncRequestEventName_NewVideoStream:            wrapFunc(api.NewVideoStream),
	event.FuncRequestEventName_NewVideoSource:            wrapFunc(api.NewVideoSource),
	event.FuncRequestEventName_CaptureVideoFrame:         wrapFunc(api.CaptureVideoFrame),
	event.FuncRequestEventName_VideoConvert:              wrapFunc(api.VideoConvert),
	event.FuncRequestEventName_VideoStreamFromParticipan: wrapFunc(api.VideoStreamFromParticipan),
}
