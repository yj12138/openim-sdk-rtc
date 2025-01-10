package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/proto/go/track"
)

func CreateVideoTrack(request *track.CreateVideoTrackRequest) *track.CreateVideoTrackResponse {
	return nil
}
func CreateAudioTrack(request *track.CreateAudioTrackRequest) *track.CreateAudioTrackResponse {
	lksdk.NewLocalTrack()
	return nil
}
func GetStats(request *track.GetStatsRequest) *track.GetStatsResponse {
	return nil
}
