package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
	"log"
)

const (
	MimeTypeOpus = "audio/opus"
)

type LocalTrack struct {
	*lksdk.LocalTrack
	// MimeType    string
	VideoWidth  int
	VideoHeight int
	Souce       *AudioSource
	Publication *lksdk.LocalTrackPublication
}

func NewAudioTrack(name string, source *AudioSource) *LocalTrack {
	// mimeType := MimeTypeOpus
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
		LocalTrack: liveKitTrack,
		Souce:      source,
		// MimeType:    mimeType,
		VideoWidth:  0,
		VideoHeight: 0,
	}
	return track
}

func NewVideoTrack() *LocalTrack {
	return nil
}
