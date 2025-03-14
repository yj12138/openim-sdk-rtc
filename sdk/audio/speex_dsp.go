package audio

/*
#include <speex/speex_echo.h>
#include <speex/speex_preprocess.h>
#include<stdio.h>
#include <stdlib.h>

SpeexPreprocessState *state;

void InitPreprocess(int frameSize,int sampleRate){
    state = speex_preprocess_state_init(frameSize, sampleRate);
    int denoise = 1;
    int noiseSuppress = -25;
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_DENOISE, &denoise);
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_NOISE_SUPPRESS, &noiseSuppress);

    int i;
    i = 0;
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_AGC, &i);
    i = 80000;
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_AGC_LEVEL, &i);
    i = 0;
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_DEREVERB, &i);
    float f = 0;
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_DEREVERB_DECAY, &f);
    f = 0;
    speex_preprocess_ctl(state, SPEEX_PREPROCESS_SET_DEREVERB_LEVEL, &f);
}

void Process(spx_int16_t* dataBuf)  {
   speex_preprocess_run(state, (spx_int16_t*)(dataBuf));
}


*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"log"
	"unsafe"
)

// AECProcessor 处理音频流（回声消除 + 降噪）
type AECProcessor struct {
}

// 初始化 AEC + NS 处理器
func NewAECProcessor(frameSize, filterLength, sampleRate int) *AECProcessor {
	C.InitPreprocess(C.int(frameSize), C.int(sampleRate))
	return &AECProcessor{}
}

func BytesToInt16(data []byte) ([]int16, error) {
	buf := bytes.NewReader(data)
	int16Data := make([]int16, len(data)/2)
	err := binary.Read(buf, binary.LittleEndian, &int16Data)
	return int16Data, err
}

// Int16ToBytes 将 []int16 转换回 []byte（小端序）
func Int16ToBytes(data []int16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	return buf.Bytes()
}

func (aec *AECProcessor) Process(input []byte) []byte {
	audioFrame, err := BytesToInt16(input)
	if err != nil {
		log.Println(err.Error())
		return input
	}
	C.Process((*C.spx_int16_t)(unsafe.Pointer(&audioFrame[0])))

	return Int16ToBytes(audioFrame)
}

// 释放资源
func (aec *AECProcessor) Destroy() {

	// C.speex_echo_state_destroy(aec.echoState)
	// C.speex_preprocess_state_destroy(aec.preprocess)
}
