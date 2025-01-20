package sdk

import (
	// pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"log"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
)

type Track struct {
	name         string
	videoWidth   int
	videoHeight  int
	liveKitTrack *lksdk.LocalTrack
}

func newLivekitTrack(mine string) (*lksdk.LocalTrack, error) {
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

func NewAudioTrack() *Track {
	track, err := newLivekitTrack(webrtc.MimeTypePCMA)
	if err != nil {
		log.Println(err.Error())
	}
	return track
}

func NewVideoTrack() *Track {

}

func CreateVideoTrack(name string) {

}

func CreateAudioTrack() *lksdk.LocalTrack {

}

func GetStats() {

}
