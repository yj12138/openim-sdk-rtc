package base

import (
	"fmt"
	"sync"
	"sync/atomic"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/sdk"
)

var (
	api *API
)

func init() {
	api = NewAPI()
}

type API struct {
	handleObjMap  sync.Map
	handleCounter atomic.Uint64
}

func NewAPI() *API {
	return &API{
		handleObjMap:  sync.Map{},
		handleCounter: atomic.Uint64{},
	}
}

func (api *API) storeObj(value any) uint64 {
	handle := api.handleCounter.Add(1)
	api.handleObjMap.Store(handle, value)
	return handle
}

func (api *API) getRoom(handle uint64) *sdk.Room {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.Room); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Room type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getLocalParticipant(handle uint64) *sdk.LocalParticipant {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.LocalParticipant); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalParticipant type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getRemoteParticipant(handle uint64) *sdk.RemoteParticipant {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.RemoteParticipant); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.RemoteParticipant type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getLocalTrack(handle uint64) *sdk.LocalTrack {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.LocalTrack); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.LocalTrack type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getRemotePublication(handle uint64) *lksdk.RemoteTrackPublication {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*lksdk.RemoteTrackPublication); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not RemoteTrackPublication type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getAudioStream(handle uint64) *sdk.AudioStream {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.AudioStream); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Stream type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getAudioSource(handle uint64) *sdk.AudioSource {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.AudioSource); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.AudioSource type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}
