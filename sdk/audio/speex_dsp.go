package audio

/*
#include<stdio.h>
#include <stdlib.h>
#include <speex/speex_echo.h>
#include <speex/speex_preprocess.h>
#include <speex/speex_resampler.h>

static SpeexPreprocessState* g_preprocess_state = 0;
static SpeexEchoState* g_echo_state = 0;

int destory(){
    if (g_preprocess_state) {
        speex_preprocess_state_destroy(g_preprocess_state);
        g_preprocess_state = 0;
    }
    if (g_echo_state) {
        speex_echo_state_destroy(g_echo_state);
        g_echo_state = 0;
    }
    return 0;
}

int init_state(int frame_size,int sample_rate){
    destory();
    int filter_length = sample_rate / 1000 * 100;
    g_echo_state = speex_echo_state_init(frame_size,filter_length);
    g_preprocess_state = speex_preprocess_state_init(frame_size,sample_rate);
    int _db = 60;
    int _denose = 1;
    int _noiseSuppress = -25;
    speex_echo_ctl(g_echo_state, SPEEX_ECHO_SET_SAMPLING_RATE, &sample_rate);
    speex_preprocess_ctl(g_preprocess_state, SPEEX_PREPROCESS_SET_DENOISE, &_denose);
    speex_preprocess_ctl(g_preprocess_state, SPEEX_PREPROCESS_SET_NOISE_SUPPRESS, &_noiseSuppress);
    speex_preprocess_ctl(g_preprocess_state, SPEEX_PREPROCESS_SET_ECHO_STATE, g_echo_state);
    speex_preprocess_ctl(g_preprocess_state, SPEEX_PREPROCESS_SET_ECHO_SUPPRESS, &_db);
    return g_echo_state ? 1 : 0 && g_preprocess_state ? 1 : 0 ;
}

int echo_cancellation(short* mic,short* play,short* out){
    if(!g_echo_state || !g_preprocess_state){
        return 1;
    }
    // speex_echo_cancellation(g_echo_state,(const spx_int16_t*)mic,(const spx_int16_t*)play,(spx_int16_t*)out);
    speex_echo_playback(g_echo_state,play);
    speex_echo_capture(g_echo_state,mic,out);
    speex_preprocess_run(g_preprocess_state,out);
    return 0;
}

*/
import "C"

import (
	"errors"
	"log"
	"unsafe"
)

type SpeexDSP struct {
	frameSize  uint32
	sampleRate uint32

	resampler        *C.SpeexResamplerState
	sourceSampleRate uint32
	targetSampleRate uint32
}

func (dsp *SpeexDSP) Init() error {
	return nil
}

func (dsp *SpeexDSP) EchoCancellation(source []byte, sampleRate uint32, echo []byte) []byte {
	if source == nil || echo == nil || len(source) != len(echo) {
		return source
	}
	frame, err := bytesToInt16(source)
	if err != nil {
		log.Println(err.Error())
		return source
	}
	echoFrame, err := bytesToInt16(echo)
	if err != nil {
		log.Println(err.Error())
		return source
	}
	frameSize := uint32(len(frame))
	if dsp.frameSize != frameSize || dsp.sampleRate != sampleRate {
		suc := C.init_state(C.int(frameSize), C.int(sampleRate))
		if suc != 0 {
			log.Println("init_state failed")
			return source
		}
		dsp.frameSize = frameSize
		dsp.sampleRate = sampleRate
	}
	outFrame := make([]int16, frameSize)
	suc := C.echo_cancellation((*C.short)(unsafe.Pointer(&frame[0])), (*C.short)(unsafe.Pointer(&echoFrame[0])), (*C.short)(unsafe.Pointer(&outFrame[0])))
	if suc != 0 {
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
	C.destory()
	if dsp.resampler != nil {
		C.speex_resampler_destroy(dsp.resampler)
	}
}
