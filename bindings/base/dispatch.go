package base

import (
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	"google.golang.org/protobuf/proto"
	"log"
	"unsafe"
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

func StoreFFICall(call *FFICall) uint64 {
	return api.storeObj(call)
}
func GetFFICall(handleId uint64) *FFICall {
	return api.getFFICall(handleId)
}
func GetFFIEvent(handleId uint64) *FFIEvent {
	return api.getFFIEvent(handleId)
}
func GetCBuffer(handleId uint64) *CBuffer {
	return api.getCBuffer(handleId)
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
	api.c.EventCallback(ffiEvent)
}

func execute(req *pb_ffi.FfiRequest) *pb_ffi.FfiResponse {
	var response = &pb_ffi.FfiResponse{}
	switch v := req.Message.(type) {
	case *pb_ffi.FfiRequest_Dispose:
		response.Message = &pb_ffi.FfiResponse_Dispose{
			Dispose: api.Dispose(v.Dispose),
		}
	case *pb_ffi.FfiRequest_Connect:
		response.Message = &pb_ffi.FfiResponse_Connect{
			Connect: api.Connect(v.Connect),
		}
	case *pb_ffi.FfiRequest_Disconnect:
		response.Message = &pb_ffi.FfiResponse_Disconnect{
			Disconnect: api.Disconnect(v.Disconnect),
		}
	case *pb_ffi.FfiRequest_PublishTrack:
		response.Message = &pb_ffi.FfiResponse_PublishTrack{
			PublishTrack: api.PublishTrack(v.PublishTrack),
		}
	case *pb_ffi.FfiRequest_UnpublishTrack:
		response.Message = &pb_ffi.FfiResponse_UnpublishTrack{
			UnpublishTrack: api.UnpublishTrack(v.UnpublishTrack),
		}
	case *pb_ffi.FfiRequest_PublishData:
		response.Message = &pb_ffi.FfiResponse_PublishData{
			PublishData: api.PublishData(v.PublishData),
		}
	case *pb_ffi.FfiRequest_SetSubscribed:
		response.Message = &pb_ffi.FfiResponse_SetSubscribed{
			SetSubscribed: api.SetSubscribed(v.SetSubscribed),
		}
	case *pb_ffi.FfiRequest_SetLocalMetadata:
		response.Message = &pb_ffi.FfiResponse_SetLocalMetadata{
			SetLocalMetadata: api.SetLocalMetadata(v.SetLocalMetadata),
		}
	case *pb_ffi.FfiRequest_SetLocalName:
		response.Message = &pb_ffi.FfiResponse_SetLocalName{
			SetLocalName: api.SetLocalName(v.SetLocalName),
		}
	case *pb_ffi.FfiRequest_SetLocalAttributes:
		response.Message = &pb_ffi.FfiResponse_SetLocalAttributes{
			SetLocalAttributes: api.SetLocalAttributes(v.SetLocalAttributes),
		}
	case *pb_ffi.FfiRequest_GetSessionStats:
		response.Message = &pb_ffi.FfiResponse_GetSessionStats{
			GetSessionStats: api.GetSessionStats(v.GetSessionStats),
		}
	case *pb_ffi.FfiRequest_PublishTranscription:
		response.Message = &pb_ffi.FfiResponse_PublishTranscription{
			PublishTranscription: api.PublishTranscription(v.PublishTranscription),
		}
	case *pb_ffi.FfiRequest_PublishSipDtmf:
		response.Message = &pb_ffi.FfiResponse_PublishSipDtmf{
			PublishSipDtmf: api.PublishSipDtmf(v.PublishSipDtmf),
		}
	case *pb_ffi.FfiRequest_CreateVideoTrack:
		response.Message = &pb_ffi.FfiResponse_CreateVideoTrack{
			CreateVideoTrack: api.CreateVideoTrack(v.CreateVideoTrack),
		}
	case *pb_ffi.FfiRequest_CreateAudioTrack:
		response.Message = &pb_ffi.FfiResponse_CreateAudioTrack{
			CreateAudioTrack: api.CreateAudioTrack(v.CreateAudioTrack),
		}
	case *pb_ffi.FfiRequest_LocalTrackMute:
		response.Message = &pb_ffi.FfiResponse_LocalTrackMute{
			LocalTrackMute: api.LocalTrackMute(v.LocalTrackMute),
		}
	case *pb_ffi.FfiRequest_EnableRemoteTrack:
		response.Message = &pb_ffi.FfiResponse_EnableRemoteTrack{
			EnableRemoteTrack: api.EnableRemoteTrack(v.EnableRemoteTrack),
		}
	case *pb_ffi.FfiRequest_GetStats:
		response.Message = &pb_ffi.FfiResponse_GetStats{
			GetStats: api.GetStats(v.GetStats),
		}
	case *pb_ffi.FfiRequest_SetTrackSubscriptionPermissions:
		response.Message = &pb_ffi.FfiResponse_SetTrackSubscriptionPermissions{
			SetTrackSubscriptionPermissions: api.SetTrackSubscriptionPermissions(v.SetTrackSubscriptionPermissions),
		}
	case *pb_ffi.FfiRequest_NewVideoStream:
		response.Message = &pb_ffi.FfiResponse_NewVideoStream{
			NewVideoStream: api.NewVideoStream(v.NewVideoStream),
		}
	case *pb_ffi.FfiRequest_NewVideoSource:
		response.Message = &pb_ffi.FfiResponse_NewVideoSource{
			NewVideoSource: api.NewVideoSource(v.NewVideoSource),
		}
	case *pb_ffi.FfiRequest_CaptureVideoFrame:
		response.Message = &pb_ffi.FfiResponse_CaptureVideoFrame{
			CaptureVideoFrame: api.CaptureVideoFrame(v.CaptureVideoFrame),
		}
	case *pb_ffi.FfiRequest_VideoConvert:
		response.Message = &pb_ffi.FfiResponse_VideoConvert{
			VideoConvert: api.VideoConvert(v.VideoConvert),
		}
	case *pb_ffi.FfiRequest_VideoStreamFromParticipant:
		response.Message = &pb_ffi.FfiResponse_VideoStreamFromParticipant{
			VideoStreamFromParticipant: api.VideoStreamFromParticipant(v.VideoStreamFromParticipant),
		}
	case *pb_ffi.FfiRequest_NewAudioStream:
		response.Message = &pb_ffi.FfiResponse_NewAudioStream{
			NewAudioStream: api.NewAudioStream(v.NewAudioStream),
		}
	case *pb_ffi.FfiRequest_NewAudioSource:
		response.Message = &pb_ffi.FfiResponse_NewAudioSource{
			NewAudioSource: api.NewAudioSource(v.NewAudioSource),
		}
	case *pb_ffi.FfiRequest_CaptureAudioFrame:
		response.Message = &pb_ffi.FfiResponse_CaptureAudioFrame{
			CaptureAudioFrame: api.CaptureAudioFrame(v.CaptureAudioFrame),
		}
	case *pb_ffi.FfiRequest_ClearAudioBuffer:
		response.Message = &pb_ffi.FfiResponse_ClearAudioBuffer{
			ClearAudioBuffer: api.ClearAudioBuffer(v.ClearAudioBuffer),
		}
	case *pb_ffi.FfiRequest_NewAudioResampler:
		response.Message = &pb_ffi.FfiResponse_NewAudioResampler{
			NewAudioResampler: api.NewAudioResampler(v.NewAudioResampler),
		}
	case *pb_ffi.FfiRequest_RemixAndResample:
		response.Message = &pb_ffi.FfiResponse_RemixAndResample{
			RemixAndResample: api.RemixAndResample(v.RemixAndResample),
		}
	case *pb_ffi.FfiRequest_E2Ee:
		response.Message = &pb_ffi.FfiResponse_E2Ee{
			E2Ee: api.E2Ee(v.E2Ee),
		}
	case *pb_ffi.FfiRequest_AudioStreamFromParticipant:
		response.Message = &pb_ffi.FfiResponse_AudioStreamFromParticipant{
			AudioStreamFromParticipant: api.AudioStreamFromParticipant(v.AudioStreamFromParticipant),
		}
	case *pb_ffi.FfiRequest_NewSoxResampler:
		response.Message = &pb_ffi.FfiResponse_NewSoxResampler{
			NewSoxResampler: api.NewSoxResampler(v.NewSoxResampler),
		}
	case *pb_ffi.FfiRequest_PushSoxResampler:
		response.Message = &pb_ffi.FfiResponse_PushSoxResampler{
			PushSoxResampler: api.PushSoxResampler(v.PushSoxResampler),
		}
	case *pb_ffi.FfiRequest_FlushSoxResampler:
		response.Message = &pb_ffi.FfiResponse_FlushSoxResampler{
			FlushSoxResampler: api.FlushSoxResampler(v.FlushSoxResampler),
		}
	case *pb_ffi.FfiRequest_SendChatMessage:
		response.Message = &pb_ffi.FfiResponse_SendChatMessage{
			SendChatMessage: api.SendChatMessage(v.SendChatMessage),
		}
	case *pb_ffi.FfiRequest_EditChatMessage:
		response.Message = &pb_ffi.FfiResponse_SendChatMessage{
			SendChatMessage: api.EditChatMessage(v.EditChatMessage),
		}
	case *pb_ffi.FfiRequest_PerformRpc:
		response.Message = &pb_ffi.FfiResponse_PerformRpc{
			PerformRpc: api.PerformRpc(v.PerformRpc),
		}
	case *pb_ffi.FfiRequest_RegisterRpcMethod:
		response.Message = &pb_ffi.FfiResponse_RegisterRpcMethod{
			RegisterRpcMethod: api.RegisterRpcMethod(v.RegisterRpcMethod),
		}
	case *pb_ffi.FfiRequest_UnregisterRpcMethod:
		response.Message = &pb_ffi.FfiResponse_UnregisterRpcMethod{
			UnregisterRpcMethod: api.UnregisterRpcMethod(v.UnregisterRpcMethod),
		}
	case *pb_ffi.FfiRequest_RpcMethodInvocationResponse:
		response.Message = &pb_ffi.FfiResponse_RpcMethodInvocationResponse{
			RpcMethodInvocationResponse: api.RpcMethodInvocationResponse(v.RpcMethodInvocationResponse),
		}
	case *pb_ffi.FfiRequest_EnableRemoteTrackPublication:
		response.Message = &pb_ffi.FfiResponse_EnableRemoteTrackPublication{
			EnableRemoteTrackPublication: api.EnableRemoteTrackPublication(v.EnableRemoteTrackPublication),
		}
	case *pb_ffi.FfiRequest_UpdateRemoteTrackPublicationDimension:
		response.Message = &pb_ffi.FfiResponse_UpdateRemoteTrackPublicationDimension{
			UpdateRemoteTrackPublicationDimension: api.UpdateRemoteTrackPublicationDimension(v.UpdateRemoteTrackPublicationDimension),
		}
	case *pb_ffi.FfiRequest_SendStreamHeader:
		response.Message = &pb_ffi.FfiResponse_SendStreamHeader{
			SendStreamHeader: api.SendStreamHeader(v.SendStreamHeader),
		}
	case *pb_ffi.FfiRequest_SendStreamChunk:
		response.Message = &pb_ffi.FfiResponse_SendStreamChunk{
			SendStreamChunk: api.SendStreamChunk(v.SendStreamChunk),
		}
	case *pb_ffi.FfiRequest_SendStreamTrailer:
		response.Message = &pb_ffi.FfiResponse_SendStreamTrailer{
			SendStreamTrailer: api.SendStreamTrailer(v.SendStreamTrailer),
		}
	case *pb_ffi.FfiRequest_SetDataChannelBufferedAmountLowThreshold:
		response.Message = &pb_ffi.FfiResponse_SetDataChannelBufferedAmountLowThreshold{
			SetDataChannelBufferedAmountLowThreshold: api.SetDataChannelBufferedAmountLowThreshold(v.SetDataChannelBufferedAmountLowThreshold),
		}
	case *pb_ffi.FfiRequest_LoadAudioFilterPlugin:
		response.Message = &pb_ffi.FfiResponse_LoadAudioFilterPlugin{
			LoadAudioFilterPlugin: api.LoadAudioFilterPlugin(v.LoadAudioFilterPlugin),
		}
	default:
	}
	return response
}

func Request(call *FFICall) {
	var req pb_ffi.FfiRequest
	call.Id = api.storeObj(call)
	err := proto.Unmarshal(call.RequestData, &req)
	if err != nil {
		log.Println("unmarshal error:", err.Error())
		return
	}
	response := execute(&req)
	if response == nil {
		log.Panic("Response is  nil")
		return
	}
	responseData, err := proto.Marshal(response)
	if err != nil {
		log.Panic("Marshal Error:", err.Error())
		return
	}
	call.ResponseData = responseData
}
