package audio

import (
	"encoding/binary"
	"log"
	"os"
)

type WavRecorder struct {
	outputName     string
	sampleRate     uint32
	numChannels    uint32
	bytesPerSample uint32

	file *os.File

	dataSize uint32
}

func (r *WavRecorder) writeHeader() {
	r.file.Write([]byte("RIFF"))                                     // RIFF 头
	binary.Write(r.file, binary.LittleEndian, uint32(36+r.dataSize)) // 总大小
	r.file.Write([]byte("WAVE"))                                     // WAVE 头
	r.file.Write([]byte("fmt "))                                     // fmt chunk
	binary.Write(r.file, binary.LittleEndian, uint32(16))            // fmt chunk 长度
	binary.Write(r.file, binary.LittleEndian, uint16(1))             // PCM 格式
	binary.Write(r.file, binary.LittleEndian, uint16(r.numChannels)) // 声道数
	binary.Write(r.file, binary.LittleEndian, uint32(r.sampleRate))  // 采样率
	byteRate := r.sampleRate * r.numChannels * r.bytesPerSample
	binary.Write(r.file, binary.LittleEndian, uint32(byteRate)) // 字节率
	blockAlign := r.numChannels * r.bytesPerSample
	binary.Write(r.file, binary.LittleEndian, uint16(blockAlign))         // 块对齐
	binary.Write(r.file, binary.LittleEndian, uint16(r.bytesPerSample*8)) // 采样位数
	r.file.Write([]byte("data"))                                          // data 头
	binary.Write(r.file, binary.LittleEndian, uint32(r.dataSize))         // 数据大小
}
func (r *WavRecorder) updateWavHeader() {
	r.file.Seek(4, 0)
	binary.Write(r.file, binary.LittleEndian, uint32(36+r.dataSize)) // 修正 RIFF chunk size
	r.file.Seek(40, 0)
	binary.Write(r.file, binary.LittleEndian, uint32(r.dataSize))
	err := r.file.Close()
	if err != nil {
		log.Println("Save Wav Err :", r.outputName, err.Error())
	}
}

func (r *WavRecorder) Start() error {
	file, err := os.Create(r.outputName)
	if err != nil {
		return err
	}
	r.file = file
	r.writeHeader()
	return nil
}

func (r *WavRecorder) Write(data []byte, end bool) {
	r.file.Write(data)
	r.dataSize += uint32(len(data) / int(r.bytesPerSample))
	if end {
		r.updateWavHeader()
	}
}

func NewWavRecorder(outputName string, sampleRate uint32, bytesPerSample uint32) *WavRecorder {
	return &WavRecorder{
		outputName:     outputName,
		sampleRate:     sampleRate,
		bytesPerSample: bytesPerSample,
	}
}
