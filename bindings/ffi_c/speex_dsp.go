package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L. -lspeexdsp
#include <speex/speex_echo.h>
#include <speex/speex_preprocess.h>
#include <stdlib.h>
*/
import "C"
import (
	"log"
	"unsafe"
)

// AECProcessor 处理音频流（回声消除 + 降噪）
type AECProcessor struct {
	echoState  *C.SpeexEchoState
	preprocess *C.SpeexPreprocessState
	frameSize  int
	sampleRate int
}

// 初始化 AEC + NS 处理器
func NewAECProcessor(frameSize, filterLength, sampleRate int) *AECProcessor {
	echoState := C.speex_echo_state_init(C.int(frameSize), C.int(filterLength))
	if echoState == nil {
		log.Fatal("Failed to initialize Speex Echo Canceller")
	}
	C.speex_echo_ctl(echoState, C.SPEEX_ECHO_SET_SAMPLING_RATE, unsafe.Pointer(&sampleRate))

	// 噪声抑制
	preprocess := C.speex_preprocess_state_init(C.int(frameSize), C.int(sampleRate))

	denoise := C.int(1)
	C.speex_preprocess_ctl(preprocess, C.SPEEX_PREPROCESS_SET_DENOISE, unsafe.Pointer(&denoise))

	return &AECProcessor{
		echoState:  echoState,
		preprocess: preprocess,
		frameSize:  frameSize,
		sampleRate: sampleRate,
	}
}

// 处理一帧音频
func (aec *AECProcessor) Process(input, echo []byte) []byte {
	output := make([]byte, len(input))

	// 指针转换
	inputPtr := (*C.spx_int16_t)(unsafe.Pointer(&input[0]))
	echoPtr := (*C.spx_int16_t)(unsafe.Pointer(&echo[0]))
	outputPtr := (*C.spx_int16_t)(unsafe.Pointer(&output[0]))

	// 回声消除
	C.speex_echo_cancel(aec.echoState, inputPtr, echoPtr, outputPtr, nil)

	// 噪声抑制
	C.speex_preprocess_run(aec.preprocess, outputPtr)

	return output
}

// 释放资源
func (aec *AECProcessor) Destroy() {
	C.speex_echo_state_destroy(aec.echoState)
	C.speex_preprocess_state_destroy(aec.preprocess)
}
