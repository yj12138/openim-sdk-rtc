package base

import (
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	pb_video "github.com/openimsdk/openim-rtc/proto/go/video_frame"
	"github.com/openimsdk/openim-rtc/sdk"
	// "log"
)

type VideoFrameBuffer struct {
	Type       pb_video.VideoBufferType
	Width      uint32
	Height     uint32
	DataPtr    uint64
	Stride     uint32
	Components []*pb_video.VideoBufferInfo_ComponentInfo
}

// Video
func (api *API) NewVideoStream(req *pb_video.NewVideoStreamRequest) (*pb_video.NewVideoStreamResponse, error) {
	track := api.getRemoteTrack(req.TrackHandle)
	videoStream := sdk.NewVideoStreamByTrack(track, req.Type, req.Format, req.NormalizeStride)
	streamHandle := api.storeObj(videoStream)

	videoStream.CallBack = func(videoFrame *sdk.VideoFrame) {
		buffer := &VideoFrameBuffer{
			Type:    videoFrame.Type,
			Width:   videoFrame.Width,
			Height:  videoFrame.Height,
			DataPtr: goByteSliceToCPointerNoCopyFunc(videoFrame.Payload),
			Stride:  0,
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_VideoStreamEvent{
				VideoStreamEvent: &pb_video.VideoStreamEvent{
					StreamHandle: streamHandle,
					Message: &pb_video.VideoStreamEvent_FrameReceived{
						FrameReceived: &pb_video.VideoFrameReceived{
							Buffer: &pb_video.OwnedVideoBuffer{
								Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(buffer)},
								Info: &pb_video.VideoBufferInfo{
									Type:    buffer.Type,
									Width:   buffer.Height,
									DataPtr: buffer.DataPtr,
									Stride:  0,
								},
							},
							TimestampUs: videoFrame.TimeStampUs,
							Rotation:    videoFrame.Rotation,
						},
					},
				},
			},
		})
	}
	res := &pb_video.NewVideoStreamResponse{
		Stream: &pb_video.OwnedVideoStream{
			Handle: &pb_handle.FfiOwnedHandle{Id: streamHandle},
			Info: &pb_video.VideoStreamInfo{
				Type: videoStream.StreamType,
			},
		},
	}
	return res, nil
}
func (api *API) NewVideoSource(req *pb_video.NewVideoSourceRequest) (*pb_video.NewVideoSourceResponse, error) {
	videoSource := sdk.NewVideoSource(req.Type, req.Resolution.Width, req.Resolution.Height)
	res := &pb_video.NewVideoSourceResponse{
		Source: &pb_video.OwnedVideoSource{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(videoSource)},
			Info: &pb_video.VideoSourceInfo{
				Type: videoSource.SourceType,
			},
		},
	}
	return res, nil
}
func (api *API) CaptureVideoFrame(req *pb_video.CaptureVideoFrameRequest) (*pb_video.CaptureVideoFrameResponse, error) {
	videoSource := api.getVideoSource(req.SourceHandle)
	// TODO 计算大小
	length := 0
	data := cPointerToGoByteSliceNoCopyFunc(req.Buffer.DataPtr, uint64(length))
	videoSource.CaptureFrame(req.TimestampUs, req.Rotation, data)
	res := &pb_video.CaptureVideoFrameResponse{}
	return res, nil
}
func (api *API) VideoConvert(req *pb_video.VideoConvertRequest) (*pb_video.VideoConvertResponse, error) {
	// TODO
	return nil, nil
}
func (api *API) VideoStreamFromParticipan(req *pb_video.VideoStreamFromParticipantRequest) (*pb_video.VideoStreamFromParticipantResponse, error) {
	// TODO
	return nil, nil
}
