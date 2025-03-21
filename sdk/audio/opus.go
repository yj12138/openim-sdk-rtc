package audio

/*
#include<opus/opus.h>
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

type OpusEncoder struct {
	encoder         *C.OpusEncoder
	encodeDataCache []byte
}

func (e *OpusEncoder) Encode(rawData []byte) ([]byte, error) {
	frame, err := bytesToInt16(rawData)
	if err != nil {
		return nil, err
	}
	frameSize := len(frame)
	maxDataBytes := len(e.encodeDataCache)
	encodedBytes := C.opus_encode(e.encoder, (*C.short)(unsafe.Pointer(&frame[0])), C.int(frameSize), (*C.uchar)(unsafe.Pointer(&e.encodeDataCache[0])), C.int(maxDataBytes))
	if encodedBytes < 0 {
		return nil, fmt.Errorf("Encoding failed:%s", C.GoString(C.opus_strerror(encodedBytes)))
	}
	encodeData := make([]byte, encodedBytes)
	copy(encodeData, e.encodeDataCache)
	return encodeData, nil
}

func (e *OpusEncoder) Destroy() {
	C.opus_encoder_destroy(e.encoder)
}

func NewOpusEncoder(sampleRate uint32, channels uint32, maxPackageBytes uint32) *OpusEncoder {
	var err C.int
	encoder := C.opus_encoder_create(C.int(sampleRate), C.int(channels), C.OPUS_APPLICATION_AUDIO, &err)
	if err != C.OPUS_OK {
		log.Println("Error create encoder:", C.GoString(C.opus_strerror(err)))
		return nil
	}
	return &OpusEncoder{
		encoder:         encoder,
		encodeDataCache: make([]byte, maxPackageBytes),
	}
}

type OpusDecoder struct {
	decoder *C.OpusDecoder
}

func (d *OpusDecoder) Decode(encodedData []byte, frameSize int32) ([]byte, error) {
	frame := make([]int16, frameSize)
	decoded_samples := C.opus_decode(d.decoder, (*C.uchar)(unsafe.Pointer(&encodedData[0])), C.int(len(encodedData)), (*C.short)(unsafe.Pointer(&frame[0])), C.int(frameSize), C.int(0))
	if decoded_samples < 0 {
		return nil, fmt.Errorf("Encoding failed:%s", C.GoString(C.opus_strerror(decoded_samples)))
	}
	return int16ToBytes(frame), nil
}

func (d *OpusDecoder) Destroy() {
	C.opus_decoder_destroy(d.decoder)
}

func NewOpusDecoder(sampleRate int32, channels int32) *OpusDecoder {
	var err C.int
	decoder := C.opus_decoder_create(C.int(sampleRate), C.int(channels), &err)
	if err != C.OPUS_OK {
		log.Println("Error create encoder:", C.GoString(C.opus_strerror(err)))
		return nil
	}
	return &OpusDecoder{
		decoder: decoder,
	}
}
