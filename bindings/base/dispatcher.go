package base

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/openimsdk/openim-rtc/proto/go/event"
	"google.golang.org/protobuf/proto"
)

var (
	handleCounter        atomic.Int64
	dispatchFfiResultFun func(handleId int64, data []byte)
)

type callFunc func(handleId int64, name event.FuncRequestEventName, req []byte) ([]byte, error)

func SetDispatchFfiResultFunc(f func(handleId int64, data []byte)) {
	dispatchFfiResultFun = f
}

func activeErrResp(handleId int64, eventName event.FuncRequestEventName, err Error) {
	var ffiResult event.FfiResult
	ffiResult.HandleId = handleId
	ffiResult.EventName = eventName
	ffiResult.ErrCode = err.Code
	ffiResult.ErrMsg = err.ErrMsg
	dispatchFfiResult(handleId, &ffiResult)
}

func activeSuccessResp(handleId int64, eventName event.FuncRequestEventName, res []byte) {
	var ffiResponse event.FfiResult
	ffiResponse.Data = res
	ffiResponse.EventName = eventName
	ffiResponse.HandleId = handleId
	dispatchFfiResult(handleId, &ffiResponse)
}

func GenerateHandleID() int64 {
	return handleCounter.Add(-1)
}

func dispatchEventResp(eventName event.FuncRequestEventName, data proto.Message) {
	var res event.FfiResult
	var err error

	res.Data, err = proto.Marshal(data)
	if err != nil {
		// TODO
		res.ErrCode = 0
		res.ErrMsg = "data marshal error"
	}

	log.Printf("PassiveEventRes")

	res.EventName = eventName
	res.HandleId = GenerateHandleID()
	dispatchFfiResult(res.HandleId, &res)
}

func dispatchFfiResult(handleId int64, result *event.FfiResult) {
	data, err := proto.Marshal(result)
	if err != nil {
		log.Panicln("[dispatchFfiResult]: marshal error: ", err.Error())
		return
	}
	if dispatchFfiResultFun != nil {
		dispatchFfiResultFun(handleId, data)
	} else {
		log.Panicln("dispatchFfiResultFun is nil")
	}
}

func SetProtocolType(protocolType int) {
	// TODO
}

func FfiRequest(data []byte) {
	go func() {
		var req event.FfiRequest
		err := proto.Unmarshal(data, &req)
		if err != nil {
			activeErrResp(0, req.EventName, Err_Proto_Marshal.New(err.Error()))
		}
		handleId := req.GetHandleId()
		if fn, ok := funcMap[req.EventName]; ok {
			res, err := fn(handleId, req.EventName, req.Data)
			if err != nil {
				activeErrResp(handleId, req.EventName, Err_SDK_Internal.New(err.Error()))
			}
			activeSuccessResp(handleId, req.EventName, res)
		} else {
			activeErrResp(handleId, req.EventName, Err_SDK_Internal.New(fmt.Sprintf("func %s not registered", req.EventName.String())))
		}
	}()
}
