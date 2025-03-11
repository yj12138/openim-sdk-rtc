package sdk

import (
	"bytes"
	"encoding/binary"
	"github.com/faiface/beep"
	"io"
)

// PCMStreamer 是一个实现了 beep.Streamer 接口的结构体，用于播放 PCM 数据
type PCMStreamer struct {
	reader      io.Reader
	sampleRate  beep.SampleRate
	numChannels int
	bitDepth    int
}

// Stream 从 PCM 数据流中读取样本，并填充到 samples 切片中
func (p *PCMStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	bytesPerSample := p.bitDepth / 8
	buffer := make([]byte, bytesPerSample*p.numChannels*len(samples))
	bytesRead, err := p.reader.Read(buffer)
	if err != nil && err != io.EOF {
		return 0, false
	}
	if bytesRead == 0 {
		return 0, false
	}

	for i := 0; i < bytesRead/(bytesPerSample*p.numChannels); i++ {
		for ch := 0; ch < p.numChannels; ch++ {
			var sample int16
			offset := (i*p.numChannels + ch) * bytesPerSample
			binary.Read(bytes.NewReader(buffer[offset:offset+bytesPerSample]), binary.LittleEndian, &sample)
			samples[i][ch] = float64(sample) / (1 << 15)
		}
	}
	return bytesRead / (bytesPerSample * p.numChannels), true
}

// Err 返回流中的错误
func (p *PCMStreamer) Err() error {
	return nil
}

type AudioResampler struct {
}

func StreamerToBytes(streamer beep.Streamer, sampleRate beep.SampleRate, numChannels int) ([]byte, error) {
	var buf bytes.Buffer
	samples := make([][2]float64, 1024)
	for {
		n, ok := streamer.Stream(samples)
		if !ok {
			break
		}
		for i := 0; i < n; i++ {
			for c := 0; c < numChannels; c++ {
				// Convert the sample to 16-bit PCM
				sample := int16(samples[i][c] * (1 << 15))
				if err := binary.Write(&buf, binary.LittleEndian, sample); err != nil {
					return nil, err
				}
			}
		}
	}
	return buf.Bytes(), nil
}

// RemixAndResample 重新采样音频数据
func (resampler *AudioResampler) RemixAndResample(pcmData []byte, sourceSampleRate uint32, sourceNumChannels int, targetSampleRate uint32) ([]byte, error) {
	oldSampleRate := beep.SampleRate(sourceSampleRate)
	newSampleRate := beep.SampleRate(targetSampleRate)
	streamer := &PCMStreamer{
		reader:      bytes.NewReader(pcmData),
		sampleRate:  oldSampleRate,
		numChannels: sourceNumChannels,
		bitDepth:    16,
	}
	resample := beep.Resample(3, oldSampleRate, newSampleRate, streamer)
	audioData, err := StreamerToBytes(resample, newSampleRate, 1)
	return audioData, err
}

func NewAudioResampler() *AudioResampler {
	return &AudioResampler{}
}
