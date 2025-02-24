package base

import (
	"fmt"
	// lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/sdk"
	"sync"
	"sync/atomic"
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

func (api *API) getParticipant(handle uint64) *sdk.Participant {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.Participant); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Participant type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}

func (api *API) getTrack(handle uint64) *sdk.Track {
	if value, ok := api.handleObjMap.Load(handle); ok {
		if r, ok := value.(*sdk.Track); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not sdk.Track type", handle))
		}
	}
	panic(fmt.Sprintf("not find handle:%d", handle))
}
