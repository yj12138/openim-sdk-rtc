package sdk

import "github.com/openimsdk/openim-rtc/sdk/audio"

func Dispose() {
	audio.GetAudioDSP().Destory()
}
