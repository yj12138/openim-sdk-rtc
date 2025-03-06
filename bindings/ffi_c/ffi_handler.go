package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

typedef void (*CallBack)(void* dataPtr,int len);
extern CallBack eventCallBack;
extern void InvokeCallBack(CallBack cb,char* dataPtr,int dataLength);
*/
import "C"

import (
	"context"
	"github.com/openimsdk/openim-rtc/bindings/base"
	"github.com/openimsdk/tools/errs"
	"github.com/openimsdk/tools/log"
	"unsafe"
)

func init() {
	base.SetEventCallBackFunc(eventCallBack)
	base.SetCPointerToGoByteSliceNoCopyFunc(CPointerToGoByteSliceNoCopy)
}

func eventCallBack(event *base.FFIEvent) {
	dataLen := len(event.Data)
	cData := (*C.uint8_t)(C.malloc(C.size_t(dataLen)))
	if cData == nil {
		log.ZWarn(context.Background(), "callback data", errs.New("Failed to allocate memory"))
	}
	cDataPtr := (*[1 << 30]byte)(unsafe.Pointer(cData))[:dataLen:dataLen]
	copy(cDataPtr, event.Data)
	if C.eventCallBack != nil {
		len := C.int(dataLen)
		C.InvokeCallBack(C.eventCallBack, (*C.char)(unsafe.Pointer(cData)), len)
	}
}

func CPointerToGoByteSliceNoCopy(cPointer uint64, length uint32) []byte {
	// 将 uint64 转换为 unsafe.Pointer
	goPointer := unsafe.Pointer(uintptr(cPointer))
	// 使用 unsafe.Slice 创建一个切片，直接引用 C 内存
	return unsafe.Slice((*byte)(goPointer), int(length))
}

//export openim_rtc_ffi_init
func openim_rtc_ffi_init(event C.CallBack) int64 {
	C.eventCallBack = event
	return 1
}

//export openim_rtc_ffi_request
func openim_rtc_ffi_request(data *C.void, length C.int, dataPtr **C.void, dataLen *C.size_t) C.int64_t {
	call := &base.FFICall{}
	// TODO 可以使用无拷贝复制
	call.ReqDataPtr = unsafe.Pointer(data)
	call.ReqData = C.GoBytes(call.ReqDataPtr, length)

	base.Request(call)

	if call.ReqData == nil {
		return 0
	}
	len := len(call.ResData)
	if len == 0 {
		return 0
	}
	ptr := C.malloc(C.size_t(len))
	if ptr == nil {
		return 0
	}
	call.ResDataPtr = unsafe.Pointer(ptr)
	C.memcpy(ptr, unsafe.Pointer(&call.ResData[0]), C.size_t(len))
	*dataPtr = (*C.void)(ptr)
	*dataLen = C.size_t(len)
	return C.int64_t(call.HandleId)
}

//export openim_rtc_ffi_drop_handle
func openim_rtc_ffi_drop_handle(handleId uint64) {
	call := base.GetFFICall(handleId)
	if call != nil {
		C.free(unsafe.Pointer(call.ReqDataPtr))
		C.free(unsafe.Pointer(call.ResDataPtr))
	}
	event := base.GetFFIEvent(handleId)
	if event != nil {
		C.free(unsafe.Pointer(event.DataPtr))
	}
	base.RemoteHandle(handleId)
}
