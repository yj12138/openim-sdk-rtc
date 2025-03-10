package base

import (
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	"google.golang.org/protobuf/proto"
	"log"
	"unsafe"
)

var (
	eventCallback func(*FFIEvent)
)

type FFICall struct {
	Id              uint64
	RequestData     []byte
	ResponseData    []byte
	RequestDataPtr  unsafe.Pointer
	ResponseDataPtr unsafe.Pointer
}

type FFIEvent struct {
	Id      uint64
	Data    []byte
	DataPtr unsafe.Pointer
}

func SetEventCallBackFunc(f func(*FFIEvent)) {
	eventCallback = f
}

func RemoteHandle(handle uint64) {
	api.delObj(handle)
}

func dispatchEvent(pbFfiEvent *pb_ffi.FfiEvent) {
	ffiEvent := &FFIEvent{}
	data, err := proto.Marshal(pbFfiEvent)
	if err != nil {
		log.Println("marshal error:", err.Error())
		return
	}
	ffiEvent.Data = data
	ffiEvent.Id = api.storeObj(ffiEvent)
	if eventCallback == nil {
		panic("eventCallback is nil")
	}
	eventCallback(ffiEvent)
}

func execute(req *pb_ffi.FfiRequest) (proto.Message, error) {
	var response proto.Message = nil
	var err error = nil
	switch v := req.Message.(type) {
	case *pb_ffi.FfiRequest_Dispose:
		response, err = api.Dispose(v.Dispose)
	case *pb_ffi.FfiRequest_Connect:
		response, err = api.Connect(v.Connect)
	case *pb_ffi.FfiRequest_Disconnect:
		response, err = api.Disconnect(v.Disconnect)
	case *pb_ffi.FfiRequest_PublishTrack:
		response, err = api.PublishTrack(v.PublishTrack)
	case *pb_ffi.FfiRequest_UnpublishTrack:
		response, err = api.UnpublishTrack(v.UnpublishTrack)
	case *pb_ffi.FfiRequest_PublishData:
		response, err = api.PublishData(v.PublishData)
	case *pb_ffi.FfiRequest_SetSubscribed:
		response, err = api.SetSubscribed(v.SetSubscribed)
	case *pb_ffi.FfiRequest_SetLocalMetadata:
		response, err = api.SetLocalMetadata(v.SetLocalMetadata)
	case *pb_ffi.FfiRequest_SetLocalName:
		response, err = api.SetLocalName(v.SetLocalName)
	case *pb_ffi.FfiRequest_SetLocalAttributes:
		response, err = api.SetLocalAttributes(v.SetLocalAttributes)
	case *pb_ffi.FfiRequest_GetSessionStats:
		response, err = api.GetSessionStats(v.GetSessionStats)
	case *pb_ffi.FfiRequest_PublishTranscription:
		response, err = api.PublishTranscription(v.PublishTranscription)
	case *pb_ffi.FfiRequest_PublishSipDtmf:
		response, err = api.PublishSipDtmf(v.PublishSipDtmf)
	case *pb_ffi.FfiRequest_CreateVideoTrack:
		response, err = api.CreateVideoTrack(v.CreateVideoTrack)
	case *pb_ffi.FfiRequest_CreateAudioTrack:
		response, err = api.CreateAudioTrack(v.CreateAudioTrack)
	case *pb_ffi.FfiRequest_LocalTrackMute:
		response, err = api.LocalTrackMute(v.LocalTrackMute)
	case *pb_ffi.FfiRequest_EnableRemoteTrack:
		response, err = api.EnableRemoteTrack(v.EnableRemoteTrack)
	case *pb_ffi.FfiRequest_GetStats:
		response, err = api.GetStats(v.GetStats)
	case *pb_ffi.FfiRequest_SetTrackSubscriptionPermissions:
		response, err = api.SetTrackSubscriptionPermissions(v.SetTrackSubscriptionPermissions)
	case *pb_ffi.FfiRequest_NewVideoStream:
		response, err = api.NewVideoStream(v.NewVideoStream)
	case *pb_ffi.FfiRequest_NewVideoSource:
		response, err = api.NewVideoSource(v.NewVideoSource)
	case *pb_ffi.FfiRequest_CaptureVideoFrame:
		response, err = api.CaptureVideoFrame(v.CaptureVideoFrame)
	case *pb_ffi.FfiRequest_VideoConvert:
		response, err = api.VideoConvert(v.VideoConvert)
	case *pb_ffi.FfiRequest_VideoStreamFromParticipant:
		response, err = api.VideoStreamFromParticipan(v.VideoStreamFromParticipant)
	case *pb_ffi.FfiRequest_NewAudioStream:
		response, err = api.NewAudioStream(v.NewAudioStream)
	case *pb_ffi.FfiRequest_NewAudioSource:
		response, err = api.NewAudioSource(v.NewAudioSource)
	case *pb_ffi.FfiRequest_CaptureAudioFrame:
		response, err = api.CaptureAudioFrame(v.CaptureAudioFrame)
	case *pb_ffi.FfiRequest_ClearAudioBuffer:
		response, err = api.ClearAudioBuffer(v.ClearAudioBuffer)
	case *pb_ffi.FfiRequest_NewAudioResampler:
		response, err = api.NewAudioResampler(v.NewAudioResampler)
	case *pb_ffi.FfiRequest_RemixAndResample:
		response, err = api.RemixAndResample(v.RemixAndResample)
	case *pb_ffi.FfiRequest_E2Ee:
		response, err = api.E2Ee(v.E2Ee)
	case *pb_ffi.FfiRequest_AudioStreamFromParticipant:
		response, err = api.AudioStreamFromParticipant(v.AudioStreamFromParticipant)
	case *pb_ffi.FfiRequest_NewSoxResampler:
		response, err = api.NewSoxResampler(v.NewSoxResampler)
	case *pb_ffi.FfiRequest_PushSoxResampler:
		response, err = api.PushSoxResampler(v.PushSoxResampler)
	case *pb_ffi.FfiRequest_FlushSoxResampler:
		response, err = api.FlushSoxResampler(v.FlushSoxResampler)
	case *pb_ffi.FfiRequest_SendChatMessage:
		response, err = api.SendChatMessage(v.SendChatMessage)
	case *pb_ffi.FfiRequest_EditChatMessage:
		response, err = api.EditChatMessage(v.EditChatMessage)
	case *pb_ffi.FfiRequest_PerformRpc:
		response, err = api.PerformRpc(v.PerformRpc)
	case *pb_ffi.FfiRequest_RegisterRpcMethod:
		response, err = api.RegisterRpcMethod(v.RegisterRpcMethod)
	case *pb_ffi.FfiRequest_UnregisterRpcMethod:
		response, err = api.UnregisterRpcMethod(v.UnregisterRpcMethod)
	case *pb_ffi.FfiRequest_RpcMethodInvocationResponse:
		response, err = api.RpcMethodInvocationResponse(v.RpcMethodInvocationResponse)
	case *pb_ffi.FfiRequest_EnableRemoteTrackPublication:
		response, err = api.EnableRemoteTrackPublication(v.EnableRemoteTrackPublication)
	case *pb_ffi.FfiRequest_UpdateRemoteTrackPublicationDimension:
		response, err = api.UpdateRemoteTrackPublicationDimension(v.UpdateRemoteTrackPublicationDimension)
	case *pb_ffi.FfiRequest_SendStreamHeader:
		response, err = api.SendStreamHeader(v.SendStreamHeader)
	case *pb_ffi.FfiRequest_SendStreamChunk:
		response, err = api.SendStreamChunk(v.SendStreamChunk)
	case *pb_ffi.FfiRequest_SendStreamTrailer:
		response, err = api.SendStreamTrailer(v.SendStreamTrailer)
	case *pb_ffi.FfiRequest_SetDataChannelBufferedAmountLowThreshold:
		response, err = api.SetDataChannelBufferedAmountLowThreshold(v.SetDataChannelBufferedAmountLowThreshold)
	case *pb_ffi.FfiRequest_LoadAudioFilterPlugin:
		response, err = api.LoadAudioFilterPlugin(v.LoadAudioFilterPlugin)
	default:
	}
	return response, err
}

func Request(call *FFICall) {
	var req pb_ffi.FfiRequest
	err := proto.Unmarshal(call.RequestData, &req)
	if err != nil {
		log.Println("unmarshal error:", err.Error())
		return
	}
	response, err := execute(&req)
	if err != nil {
		log.Println("Request Error:", err.Error())
		return
	}
	if response == nil {
		log.Println("Response is  nil")
		return
	}
	responseData, err := proto.Marshal(response)
	if err != nil {
		log.Println("Marshal Error:", err.Error())
		return
	}
	call.ResponseData = responseData
}
