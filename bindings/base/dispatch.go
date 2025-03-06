package base

import (
	"fmt"
	"log"
	"unsafe"

	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	"google.golang.org/protobuf/proto"
)

var (
	eventCallback func(*FFIEvent)
)

type FFICall struct {
	HandleId   uint64
	ReqData    []byte
	ResData    []byte
	ReqDataPtr unsafe.Pointer
	ResDataPtr unsafe.Pointer
}

type FFIEvent struct {
	HandleId uint64
	Data     []byte
	DataPtr  unsafe.Pointer
}

func SetEventCallBackFunc(f func(*FFIEvent)) {
	eventCallback = f
}

func (a *API) getFFICall(handle uint64) *FFICall {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*FFICall); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not FFICall type", handle))
		}
	}
	return nil
}
func (a *API) getFFIEvent(handle uint64) *FFIEvent {
	if value, ok := api.objMap.Load(handle); ok {
		if r, ok := value.(*FFIEvent); ok {
			return r
		} else {
			panic(fmt.Sprintf("handle:%d is not FFIEvent type", handle))
		}
	}
	return nil
}
func StoreFFICall(call *FFICall) uint64 {
	return api.storeObj(call)
}
func GetFFICall(handle uint64) *FFICall {
	return api.getFFICall(handle)
}
func GetFFIEvent(handle uint64) *FFIEvent {
	return api.getFFIEvent(handle)
}

func RemoteHandle(handle uint64) {
	api.delObj(handle)
}

func dispatchEvent(pbFfiEvent *pb_ffi.FfiEvent) {
	ffiEvent := &FFIEvent{}
	data, err := proto.Marshal(pbFfiEvent)
	if err != nil {
		ffiEvent.Data = data
		ffiEvent.HandleId = api.storeObj(ffiEvent)
		eventCallback(ffiEvent)
	} else {
		log.Println("marshal error:", err.Error())
	}
}

func Request(call *FFICall) {
	var req pb_ffi.FfiRequest
	err := proto.Unmarshal(call.ReqData, &req)
	if err != nil {
		log.Println("unmarshal error:", err.Error())
		return
	}
    
}
