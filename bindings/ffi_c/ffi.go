package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

typedef void (*CallBack)(void* dataPtr,int len);
extern CallBack eventCallBack;
extern void InvokeCallBack(CallBack cb,char* dataPtr,int dataLength);
extern void CPrint(char* dataPtr,int dataLength);
*/
import "C"

import (
	"log"
	"unsafe"

	"github.com/openimsdk/openim-rtc/bindings/base"
)

func init() {
	log.SetFlags(log.Llongfile)
	base.InitAPI(newCInterface())
}

//export openim_rtc_ffi_init
func openim_rtc_ffi_init(event C.CallBack, captureLogs C.int, sdk *C.char, sdkVersion *C.char) int64 {
	log.Println(captureLogs, sdk, sdkVersion)
	C.eventCallBack = event
	return 1
}

//export openim_rtc_ffi_request
func openim_rtc_ffi_request(data *C.void, length C.int, dataPtr **C.uint8_t, dataLen *C.uintptr_t) C.int64_t {
	call := &base.FFICall{}
	// TODO 可以使用无拷贝复制
	call.RequestDataPtr = unsafe.Pointer(data)
	call.RequestData = C.GoBytes(call.RequestDataPtr, length)
	// log.Println("openim_rtc_ffi request", call.RequestData)
	base.Request(call)
	// log.Println("openim_rtc_ffi reponse ", call.ResponseData)
	if call.ResponseData == nil {
		return 0
	}
	responseData := call.ResponseData
	resLength := len(responseData)
	if resLength == 0 {
		return 0
	}
	// 分配 C 内存
	cResponse := C.malloc(C.size_t(resLength))
	if cResponse == nil {
		return 0 // 内存分配失败，返回 0 句柄
	}
	// log.Println("Go Print", responseData)
	// 复制数据到 C 内存
	copy((*[1 << 30]byte)(cResponse)[:resLength:resLength], responseData)
	call.ResponseDataPtr = cResponse
	// C.CPrint((*C.char)(cResponse), C.int(resLength))
	// 设置返回指针和长度
	*dataPtr = (*C.uint8_t)(cResponse)
	*dataLen = C.uintptr_t(resLength)
	// log.Println("openim_rtc_ffi call", call)
	return C.int64_t(call.Id)
}

//export openim_rtc_ffi_drop_handle
func openim_rtc_ffi_drop_handle(handleId uint64) {
	// log.Println("openim_rtc_ffi drop handle", handleId)
	call := base.GetFFICall(handleId)
	if call != nil {
		// log.Println("drop FFICall", call)
		// 由C端释放
		// C.free(unsafe.Pointer(call.RequestDataPtr))
		C.free(unsafe.Pointer(call.ResponseDataPtr))
	}
	event := base.GetFFIEvent(handleId)
	if event != nil {
		// log.Println("drop FFIEvent", event)
		C.free(unsafe.Pointer(event.DataPtr))
	}
	cbuffer := base.GetCBuffer(handleId)
	if cbuffer != nil {
		// 直接使用的go的数据指针不用清理内存
	}
	base.RemoteHandle(handleId)
}
