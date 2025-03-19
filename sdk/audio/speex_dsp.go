package audio

/*
#include<stdio.h>
#include <stdlib.h>
#include <speex/speex_echo.h>
#include <speex/speex_preprocess.h>
#include <speex/speex_resampler.h>
*/
import "C"

import (
	"errors"
	"log"
	"unsafe"
)

const EchoTail = 4800 // Echo cancellation delay buffer size

type SpeexDSP struct {
	echoState  *C.SpeexEchoState
	preState   *C.SpeexPreprocessState
	frameSize  uint32
	sampleRate uint32

	resampler        *C.SpeexResamplerState
	sourceSampleRate uint32
	targetSampleRate uint32
}

func (dsp *SpeexDSP) Init() error {
	return nil
}

func (dsp *SpeexDSP) initState(frameSize uint32, sampleRate uint32, useAEC bool) error {
	if dsp.preState != nil {
		C.speex_preprocess_state_destroy(dsp.preState)
	}
	preprocessState := C.speex_preprocess_state_init(C.int(frameSize), C.int(sampleRate))
	if preprocessState == nil {
		return errors.New("failed to initialize Speex preprocess")
	}
	if useAEC {
		echoState := C.speex_echo_state_init(C.int(frameSize), C.int(EchoTail))
		if echoState == nil {
			C.speex_preprocess_state_destroy(preprocessState)
			return errors.New("failed to initialize Speex echo canceller")
		}
		C.speex_echo_ctl(echoState, C.SPEEX_ECHO_SET_SAMPLING_RATE, unsafe.Pointer(&sampleRate))
		dsp.echoState = echoState
	}
	if useAEC {
		C.speex_preprocess_ctl(preprocessState, C.SPEEX_PREPROCESS_SET_ECHO_STATE, unsafe.Pointer(dsp.echoState))
	}

	dsp.preState = preprocessState
	dsp.frameSize = frameSize
	dsp.sampleRate = sampleRate
	return nil
}

func (dsp *SpeexDSP) AAAProcess(source []byte, sampleRate uint32, echo []byte) []byte {
	frame, err := bytesToInt16(source)
	if err != nil {
		log.Println("AAAProcess", err.Error())
		return source
	}
	frameSize := uint32(len(frame))
	useAEC := echo != nil || len(echo) == len(source)
	if useAEC {
		if len(source) != len(echo) {
			log.Println("AAASProcess", "Source Frame Size != Echo Frame Size", len(source), len(echo))
			return source
		}
	}
	if dsp.frameSize != frameSize || dsp.sampleRate != sampleRate {
		err := dsp.initState(frameSize, sampleRate, useAEC)
		if err != nil {
			log.Println("AAAProcess", err.Error())
			return source
		}
	}
	outFrame := frame
	if useAEC {
		echoFrame, err := bytesToInt16(echo)
		if err != nil {
			log.Println("AAAProcess", err.Error())
			return source
		}
		outFrame := make([]int16, frameSize)
		C.speex_echo_cancellation(dsp.echoState, (*C.spx_int16_t)(unsafe.Pointer(&frame[0])), (*C.spx_int16_t)(unsafe.Pointer(&echoFrame[0])), (*C.spx_int16_t)(unsafe.Pointer(&outFrame[0])))
	}
	processed := C.speex_preprocess_run(dsp.preState, (*C.spx_int16_t)(unsafe.Pointer(&outFrame[0])))
	if processed == 0 {
		log.Println("AAAProcess", "failed to process audio frame")
		return source
	}
	return int16ToBytes(outFrame)
}

func (dsp *SpeexDSP) initResampler(sourceSampleRate uint32, targetSampleRate uint32) error {
	if dsp.resampler != nil {
		C.speex_resampler_destroy(dsp.resampler)
	}
	var errCode C.int
	resampler := C.speex_resampler_init(1, C.uint(sourceSampleRate), C.uint(targetSampleRate), C.int(3), &errCode)
	if resampler == nil || errCode != 0 {
		return errors.New("failed to initialize Speex resampler")
	}
	dsp.sourceSampleRate = sourceSampleRate
	dsp.targetSampleRate = targetSampleRate
	dsp.resampler = resampler
	return nil
}

func (dsp *SpeexDSP) Resample(sourceData []byte, sourceSampleRate uint32, sourceNumChannels uint32, targetSampleRate uint32, targetNumChannels uint32) []byte {
	if sourceSampleRate != dsp.sourceSampleRate || targetSampleRate != dsp.targetSampleRate {
		err := dsp.initResampler(sourceSampleRate, targetSampleRate)
		if err != nil {
			log.Println("Resample", err.Error())
			return sourceData
		}
	}
	if sourceNumChannels != 1 && sourceNumChannels != 2 {
		log.Println("unsupported input channels")
		return sourceData
	}
	if targetNumChannels != 1 && targetNumChannels != 2 {
		log.Println("unsupported output channels")
		return sourceData
	}
	input, err := bytesToInt16(sourceData)

	outLen := len(input) / int(sourceNumChannels) * int(targetSampleRate) / int(sourceSampleRate) * int(targetNumChannels)
	output := make([]int16, outLen)

	var inLenC, outLenC C.uint
	inLenC = C.uint(len(input) / int(sourceNumChannels))
	outLenC = C.uint(outLen / int(targetNumChannels))
	if err != nil {
		log.Println(err.Error())
		return sourceData
	}
	if sourceNumChannels == 2 && targetNumChannels == 1 {
		monoInput := make([]int16, len(input)/2)
		for i := 0; i < len(monoInput); i++ {
			monoInput[i] = (input[2*i] + input[2*i+1]) / 2
		}
		input = monoInput
	} else if sourceNumChannels == 1 && targetNumChannels == 2 {
		stereoOutput := make([]int16, len(input)*2)
		for i := 0; i < len(input); i++ {
			stereoOutput[2*i] = input[i]
			stereoOutput[2*i+1] = input[i]
		}
		input = stereoOutput
	}

	res := C.speex_resampler_process_int(dsp.resampler, 0, (*C.spx_int16_t)(unsafe.Pointer(&input[0])), &inLenC, (*C.spx_int16_t)(unsafe.Pointer(&output[0])), &outLenC)
	if res != 0 {
		log.Println("failed to resample audio")
		return sourceData
	}
	outputFrame := output[:outLenC*C.uint(targetNumChannels)]
	return int16ToBytes(outputFrame)
}

func (dsp *SpeexDSP) Destory() {
	if dsp.echoState != nil {
		C.speex_echo_state_destroy(dsp.echoState)
	}
	if dsp.preState != nil {
		C.speex_preprocess_state_destroy(dsp.preState)
	}
}
