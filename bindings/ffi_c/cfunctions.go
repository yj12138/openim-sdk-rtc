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

type CFunctions struct {
}

func (c *CFunctions) EventCallback(event *base.FFIEvent) {
	dataLen := len(event.Data)
	cData := (*C.uint8_t)(C.malloc(C.size_t(dataLen)))
	if cData == nil {
		log.Println("callback data", "Failed to allocate memory")
	}
	cDataPtr := (*[1 << 30]byte)(unsafe.Pointer(cData))[:dataLen:dataLen]
	copy(cDataPtr, event.Data)
	if C.eventCallBack != nil {
		len := C.int(dataLen)
		C.InvokeCallBack(C.eventCallBack, (*C.char)(unsafe.Pointer(cData)), len)
	}
}
func (c *CFunctions) CPointerToGoByteSliceNoCopy(cPointer uint64, length uint64) []byte {
	// 将 uint64 转换为 unsafe.Pointer
	goPointer := unsafe.Pointer(uintptr(cPointer))
	// 使用 unsafe.Slice 创建一个切片，直接引用 C 内存
	return unsafe.Slice((*byte)(goPointer), int(length))
}

func (c *CFunctions) GoByteSliceToCPointerNoCopy(data []byte) uint64 {
	if len(data) == 0 {
		panic("GoByteSliceToCPointerNoCopy Length = 0")
	}
	return uint64(uintptr(unsafe.Pointer(&data[0])))
}

func (c *CFunctions) GetAudio_AEC_NS_Frame(input []byte, sampleRate uint32) []byte {
    
	return nil
}

func newCInterface() *CFunctions {
	return &CFunctions{}
}
