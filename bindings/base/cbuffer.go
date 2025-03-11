package base

type CBuffer struct {
	SourceData []byte
	CDataPtr   uint64
	CDataLen   int
}

func NewCBuffer(sourceData []byte) *CBuffer {
	ptr := goByteSliceToCPointerNoCopyFunc(sourceData)
	return &CBuffer{
		SourceData: sourceData,
		CDataPtr:   ptr,
		CDataLen:   len(sourceData),
	}
}
