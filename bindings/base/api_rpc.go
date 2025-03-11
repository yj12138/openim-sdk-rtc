package base

import (
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_rpc "github.com/openimsdk/openim-rtc/proto/go/rpc"
)

func (api *API) PerformRpc(req *pb_rpc.PerformRpcRequest) *pb_rpc.PerformRpcResponse {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		participant.PerformRpc(req.DestinationIdentity, req.Method, req.Payload, req.ResponseTimeoutMs)
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_PerformRpc{
				PerformRpc: &pb_rpc.PerformRpcCallback{
					AsyncId: asyncId,
				},
			},
		})
	}()
	return &pb_rpc.PerformRpcResponse{
		AsyncId: asyncId,
	}
}

func (api *API) RegisterRpcMethod(req *pb_rpc.RegisterRpcMethodRequest) *pb_rpc.RegisterRpcMethodResponse {
	return nil
}

func (api *API) UnregisterRpcMethod(req *pb_rpc.UnregisterRpcMethodRequest) *pb_rpc.UnregisterRpcMethodResponse {
	return nil
}

func (api *API) RpcMethodInvocationResponse(req *pb_rpc.RpcMethodInvocationResponseRequest) *pb_rpc.RpcMethodInvocationResponseResponse {
	return nil
}
