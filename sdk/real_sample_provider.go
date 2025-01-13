package sdk

import (
	"context"
	"log"
	"math"
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

const (
	sampleRate = 8000   // 采样率 8kHz（符合 G.711）
	channels   = 1      // 单声道
	duration   = 1.0    // 音频持续时间，1秒
	maxInt16   = 32767  // int16 最大值
	minInt16   = -32768 // int16 最小值
)

func generateSineWave(frequency float64, duration float64) []int16 {
	// 生成正弦波 PCM 数据
	numSamples := int(sampleRate * duration)
	data := make([]int16, numSamples)
	for i := 0; i < numSamples; i++ {
		// 使用公式生成正弦波 PCM 数据
		sample := math.Sin(2 * math.Pi * frequency * float64(i) / float64(sampleRate))
		// 将浮动的样本值转换为 16 位 PCM 数据，并确保在 int16 范围内
		data[i] = int16(sample * float64(maxInt16))
	}
	return data
}

func pcmToPCMA(pcm []int16) []byte {
	// 将 PCM 数据转换为 PCMA（A-law）编码
	encodedData := make([]byte, len(pcm)) // 假设每个 PCM 数据会变成一个字节
	for i, sample := range pcm {
		// 这里进行 PCM 到 A-law 编码（简化为示例）
		encodedData[i] = byte((sample + maxInt16) >> 8) // 简化转换
	}
	return encodedData
}
func combineChords(frequencies []float64) []int16 {
	// 合成和弦音频：将多个正弦波信号合并为一个
	numSamples := int(sampleRate * duration)
	combined := make([]int16, numSamples)
	for _, freq := range frequencies {
		sineWave := generateSineWave(freq, duration)
		for i := 0; i < numSamples; i++ {
			combined[i] += sineWave[i]
			// 防止溢出，确保合成后的值在 int16 范围内
			if combined[i] > maxInt16 {
				combined[i] = maxInt16
			}
			if combined[i] < minInt16 {
				combined[i] = minInt16
			}
		}
	}
	return combined
}

type RealSampleProvider struct {
	Mime            string
	FrameDuration   time.Duration
	OnWriteComplete func()
}

func (p *RealSampleProvider) OnBind() error {
	log.Println("RealSampleProvider OnBind")
	return nil
}

func (p *RealSampleProvider) OnUnbind() error {
	log.Println("RealSampleProvider OnUnbind")
	return nil
}

func (p *RealSampleProvider) Close() error {
	log.Println("RealSampleProvider Close")
	return nil
}

func (p RealSampleProvider) CurrentAudioLevel() uint8 {
	return 0
}

func (p *RealSampleProvider) NextSample(c context.Context) (media.Sample, error) {
	log.Println("RealSampleProvider NextSample")
	sample := media.Sample{}
	switch p.Mime {
	case webrtc.MimeTypePCMA:
		frequencies := []float64{440.0, 554.37, 659.26}
		combined := combineChords(frequencies)
		pcmaData := pcmToPCMA(combined)
		sample.Data = pcmaData
		sample.Duration = duration
	}
	return sample, nil
}

func NewRealTrack(mine string) (*lksdk.LocalTrack, error) {
	provider := &RealSampleProvider{
		Mime: mine,
	}

	track, err := lksdk.NewLocalTrack(webrtc.RTPCodecCapability{
		MimeType: provider.Mime,
	})

	if err != nil {
		return nil, err
	}

	track.OnBind(func() {
		// write audio stream
		if err := track.StartWrite(provider, provider.OnWriteComplete); err != nil {
			log.Println(err.Error())
		}
	})

	return track, nil
}
