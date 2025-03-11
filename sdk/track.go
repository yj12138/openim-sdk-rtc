package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
	"log"
)

type LocalTrack struct {
	*lksdk.LocalTrack
	// MimeType    string
	Name        string
	VideoWidth  int
	VideoHeight int
	AudioSource *AudioSource
	VideoSource *VideoSource
	Publication *lksdk.LocalTrackPublication
}

func NewAudioTrack(name string, source *AudioSource) *LocalTrack {
	liveKitTrack, err := lksdk.NewLocalTrack(webrtc.RTPCodecCapability{
		MimeType: source.mime,
	})
	if err != nil {
		log.Panic(err.Error())
	}
	liveKitTrack.OnBind(func() {
		if err := liveKitTrack.StartWrite(source, source.onWriteComplete); err != nil {
			log.Panic(err.Error())
		}
	})
	track := &LocalTrack{
		Name:        name,
		LocalTrack:  liveKitTrack,
		AudioSource: source,
		VideoWidth:  0,
		VideoHeight: 0,
	}
	return track
}

func NewVideoTrack(name string, source *VideoSource) *LocalTrack {
	liveKitTrack, err := lksdk.NewLocalTrack(webrtc.RTPCodecCapability{
		MimeType: source.mime,
	})
	if err != nil {
		log.Panic(err.Error())
	}
	liveKitTrack.OnBind(func() {
		if err := liveKitTrack.StartWrite(source, source.onWriteComplete); err != nil {
			log.Panic(err.Error())
		}
	})
	track := &LocalTrack{
		Name:        name,
		LocalTrack:  liveKitTrack,
		VideoSource: source,
		VideoWidth:  0,
		VideoHeight: 0,
	}
	return track
}
