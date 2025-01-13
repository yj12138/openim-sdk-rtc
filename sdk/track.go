package sdk

import (
	// pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"log"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4"
)

func CreateVideoTrack() {
}

func CreateAudioTrack() *lksdk.LocalTrack {
	track, err := NewRealTrack(webrtc.MimeTypePCMA)
	if err != nil {
		log.Println(err.Error())
	}
	return track
}

func GetStats() {

}
