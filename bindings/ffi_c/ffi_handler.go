package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

typedef void (*CallBack)(void* dataPtr,int len);
extern CallBack eventCallBack;
extern void InvokeCallBack(CallBack cb,char* dataPtr,int dataLength);
*/
import "C"

import (
	"context"
	"fmt"
	"github.com/openimsdk/tools/errs"
	"github.com/openimsdk/tools/log"
	"sync"
	"time"
	"unsafe"
)

type ResultData struct {
	cData  *C.uint8_t
	length int
}

var (
	resultMap = make(map[int64]*ResultData)
	mu        sync.Mutex
)

const (
	checkTimePeriod = time.Minute * 1
)

func init() {
	go monitorResultMapSize()
}

func dispatchResultForC(handleID int64, data []byte) {
	cData := (*C.uint8_t)(C.malloc(C.size_t(len(data))))
	if cData == nil {
		log.ZWarn(context.Background(), "callback data", errs.New("Failed to allocate memory"))
	}
	cDataPtr := (*[1 << 30]byte)(unsafe.Pointer(cData))[:len(data):len(data)]
	copy(cDataPtr, data)
	mu.Lock()
	resultMap[handleID] = &ResultData{
		cData:  cData,
		length: len(data),
	}
	mu.Unlock()
	if C.eventCallBack != nil {
		len := C.int(len(data))
		C.InvokeCallBack(C.eventCallBack, (*C.char)(unsafe.Pointer(cData)), len)
	}
	// Currently, the SDK uses asynchronous calls for Go to C interface
	//exports to other languages, so no return value is needed here.
}
func monitorResultMapSize() {
	ticker := time.NewTicker(checkTimePeriod)
	defer ticker.Stop()

	for range ticker.C {
		mu.Lock()
		totalBytes := 0
		for _, result := range resultMap {
			totalBytes += result.length // Use the actual length stored in ResultData
		}
		mu.Unlock()
		totalMB := float64(totalBytes) / (1024 * 1024)
		log.ZDebug(context.Background(), fmt.Sprintf("Current resultMap size: %.2f MB", totalMB))
	}
}

// ffi_init initializes the callback and a selected serialization protocol
// event: The callback function to be invoked.
//
//export ffi_init
func ffi_init(event C.CallBack) int64 {
	C.eventCallBack = event
	return 1
}

//export ffi_request
func ffi_request(data *C.void, length C.int) {
	//Synchronously copy data to prevent memory from being released prematurely after calling the Go function.
	// goData := C.GoBytes(unsafe.Pointer(data), length)
	// base.FfiRequest(goData)
}

//export ffi_drop_handle
func ffi_drop_handle(handleID int64) {
	mu.Lock()
	defer mu.Unlock()
	if result, ok := resultMap[handleID]; ok {
		C.free(unsafe.Pointer(result.cData))
		delete(resultMap, handleID)
	} else {
		log.ZWarn(context.Background(), "can not find resource to recycle", nil, "handleID", handleID)
	}
}
