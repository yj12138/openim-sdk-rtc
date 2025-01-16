package sdk

import "github.com/openimsdk/openim-rtc/proto/go/ffi"

var listener RoomListener

func Dispose(request *ffi.DisposeRequest) *ffi.DisposeResponse {
	return nil
}

func SetRoomListener(l RoomListener) {
	listener = l
}
