package base

type CInterface interface {
	EventCallback(*FFIEvent)
	CPointerToGoByteSliceNoCopy(cPointer uint64, length uint64) []byte
	GoByteSliceToCPointerNoCopy(data []byte) uint64
	GetAudio_AEC_NS_Frame(input []byte, sampleRate uint32) []byte
}
