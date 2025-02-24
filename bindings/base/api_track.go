package base

import (
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) CreateVideoTrack(req *pb_track.CreateVideoTrackReq) (*pb_track.CreateVideoTrackRes, error) {
	track := sdk.NewVideoTrack()
	trackHandle := api.storeObj(track)
	res := &pb_track.CreateVideoTrackRes{
		Track: &pb_track.OwnedTrack{
			Handle: trackHandle,
			Info: &pb_track.TrackInfo{
				Sid:         track.GetSid(),
				Name:        track.GetName(),
				Kind:        track.Kind(),
				StreamState: track.StreamState(),
				Muted:       track.IsMuted(),
				Remote:      track.IsRemote(),
			},
		},
	}
	return res, nil
}
func (api *API) CreateAudioTrack(req *pb_track.CreateAudioTrackReq) (*pb_track.CreateAudioTrackRes, error) {
	track := sdk.NewAudioTrack(req.Name)
	trackHandle := api.storeObj(track)
	res := &pb_track.CreateAudioTrackRes{
		Track: &pb_track.OwnedTrack{
			Handle: trackHandle,
			Info: &pb_track.TrackInfo{
				Sid:         track.GetSid(),
				Name:        track.GetName(),
				Kind:        track.Kind(),
				StreamState: track.StreamState(),
				Muted:       track.IsMuted(),
				Remote:      track.IsRemote(),
			},
		},
	}
	return res, nil
}
func (api *API) LocalTrackMute(req *pb_track.LocalTrackMuteReq) (*pb_track.LocalTrackMuteRes, error) {
	track := api.getTrack(req.TrackHandle)
	track.SetMuted(req.Mute)
	res := &pb_track.LocalTrackMuteRes{}
	return res, nil
}
func (api *API) EnableRemoteTrack(req *pb_track.EnableRemoteTrackReq) (*pb_track.EnableRemoteTrackRes, error) {
	return nil, nil
}
func (api *API) GetStats(req *pb_track.GetStatsReq) (*pb_track.GetStatsRes, error) {

	return nil, nil
}
func (api *API) SetTrackSubscriptionPermissions(req *pb_track.SetTrackSubscriptionPermissionsReq) (*pb_track.SetTrackSubscriptionPermissionsRes, error) {

	return nil, nil
}
