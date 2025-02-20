package base

import (
	"context"
	"github.com/openimsdk/openim-rtc/proto/go/event"
	"sync/atomic"
)

var (
	handleCounter        atomic.Int64
	dispatchFfiResultFun func(handleId int64, data []byte)
	sendFfiRequestFun    func(ctx context.Context, handleID int64, data []byte) ([]byte, error)
)

type callFunc func(handleId int64, name event.FuncRequestEventName, req []byte) ([]byte, error)

func SetDispatchFfiResultFunc(f func(handleId int64, data []byte)) {
	dispatchFfiResultFun = f
}

func SetSendFfiRequestFunc(f func(ctx context.Context, handleID int64, data []byte) ([]byte, error)) {
	sendFfiRequestFun = f
}

func SetProtocolType(protocolType int) {
	// TODO
}

func FfiRequest(data []byte) {
	// TODO
}
