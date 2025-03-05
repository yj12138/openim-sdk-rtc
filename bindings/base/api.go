package base

import (
	"fmt"
	pb_common "github.com/openimsdk/openim-rtc/proto/go/common"
	"github.com/openimsdk/openim-rtc/sdk"
	"sync"
	"sync/atomic"
)

type CPointerToGoByteSliceNoCopyFunc func(cPointer uint64, length uint32) []byte

var (
	api                             *API
	cPointerToGoByteSliceNoCopyFunc CPointerToGoByteSliceNoCopyFunc
)

func init() {
	api = NewAPI()
}

func SetCPointerToGoByteSliceNoCopyFunc(f CPointerToGoByteSliceNoCopyFunc) {
	cPointerToGoByteSliceNoCopyFunc = f
}

type API struct {
	objMap  sync.Map
	counter atomic.Uint64
}

func NewAPI() *API {
	return &API{
		objMap:  sync.Map{},
		counter: atomic.Uint64{},
	}
}

func (api *API) storeObj(value any) *pb_common.FfiOwnedHandle {
	handle := api.counter.Add(1)
	api.objMap.Store(handle, value)
	return &pb_common.FfiOwnedHandle{
		Id: handle,
	}
}

func (api *API) getRoom(handle *pb_common.FfiOwnedHandle) *sdk.Room {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.Room); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Room type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getLocalParticipant(handle *pb_common.FfiOwnedHandle) *sdk.LocalParticipant {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.LocalParticipant); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalParticipant type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getRemoteParticipant(handle *pb_common.FfiOwnedHandle) *sdk.RemoteParticipant {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.RemoteParticipant); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.RemoteParticipant type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getLocalTrack(handle *pb_common.FfiOwnedHandle) *sdk.LocalTrack {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.LocalTrack); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalTrack type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getRemoteTrack(handle *pb_common.FfiOwnedHandle) *sdk.RemoteTrack {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.RemoteTrack); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalTrack type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getAudioStream(handle *pb_common.FfiOwnedHandle) *sdk.AudioStream {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.AudioStream); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Stream type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getAudioSource(handle *pb_common.FfiOwnedHandle) *sdk.AudioSource {
	if value, ok := api.objMap.Load(handle.Id); ok {
		if r, ok := value.(*sdk.AudioSource); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.AudioSource type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}
