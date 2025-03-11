package base

import (
	"fmt"
	"sync"
	"sync/atomic"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/sdk"
	"github.com/pion/webrtc/v4"
)

type CPointerToGoByteSliceNoCopyFunc func(cPointer uint64, length uint64) []byte
type GoByteSliceToCPointerNoCopyFunc func(data []byte) uint64

var (
	api                             *API
	cPointerToGoByteSliceNoCopyFunc CPointerToGoByteSliceNoCopyFunc
	goByteSliceToCPointerNoCopyFunc GoByteSliceToCPointerNoCopyFunc
)

func init() {
	api = newAPI()
}

// C数据转go数据
func SetCPointerToGoByteSliceNoCopyFunc(f CPointerToGoByteSliceNoCopyFunc) {
	cPointerToGoByteSliceNoCopyFunc = f
}

// go数据转C数据
func SetGoByteSliceToCPointerNoCopyFunc(f GoByteSliceToCPointerNoCopyFunc) {
	goByteSliceToCPointerNoCopyFunc = f
}

type API struct {
	objMap         sync.Map
	handleCounter  atomic.Uint64
	asyncIdCounter atomic.Uint64
}

func newAPI() *API {
	return &API{
		objMap:        sync.Map{},
		handleCounter: atomic.Uint64{},
	}
}

func (api *API) nextAsyncId() uint64 {
	return api.asyncIdCounter.Add(1)
}

func (api *API) storeObj(value any) uint64 {
	handle := api.handleCounter.Add(1)
	api.objMap.Store(handle, value)
	return handle
}
func (api *API) delObj(handle uint64) {
	api.objMap.Delete(handle)
}

func (api *API) getRoom(handle uint64) *sdk.Room {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.Room); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Room type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getLocalParticipant(handle uint64) *sdk.LocalParticipant {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.LocalParticipant); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalParticipant type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getLocalTrack(handle uint64) *sdk.LocalTrack {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.LocalTrack); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalTrack type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getRemoteTrack(handle uint64) *webrtc.TrackRemote {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*webrtc.TrackRemote); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalTrack type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getRemoteTackPublication(handle uint64) *lksdk.RemoteTrackPublication {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*lksdk.RemoteTrackPublication); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not lksdk.RemoteTrackPublication type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getAudioStream(handle uint64) *sdk.AudioStream {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.AudioStream); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Stream type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getAudioSource(handle uint64) *sdk.AudioSource {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.AudioSource); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.AudioSource type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) GetAudioResampler(handle uint64) *sdk.AudioResampler {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.AudioResampler); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.AudioResampler type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getVideoStream(handle uint64) *sdk.VideoStream {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.VideoStream); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Stream type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getVideoSource(handle uint64) *sdk.VideoSource {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*sdk.VideoSource); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.VideoSource type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (a *API) getFFICall(handleId uint64) *FFICall {
	if value, ok := api.objMap.Load(handleId); ok {
		if r, ok := value.(*FFICall); ok {
			return r
		}
	}
	return nil
}
func (a *API) getFFIEvent(handleId uint64) *FFIEvent {
	if value, ok := api.objMap.Load(handleId); ok {
		if r, ok := value.(*FFIEvent); ok {
			return r
		}
	}
	return nil
}
