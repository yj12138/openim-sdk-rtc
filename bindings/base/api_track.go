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
				Sid:         track.LocalTrackPublication.SID(),
				Name:        track.LocalTrackPublication.Name(),
				Kind:        sdk.ConvertTrackKind(track.LocalTrackPublication.Kind()),
				StreamState: pb_track.StreamState_STATE_ACTIVE,
				Muted:       track.LocalTrackPublication.IsMuted(),
				Remote:      false,
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
				Sid:         track.LocalTrackPublication.SID(),
				Name:        track.LocalTrackPublication.Name(),
				Kind:        sdk.ConvertTrackKind(track.LocalTrackPublication.Kind()),
				StreamState: pb_track.StreamState_STATE_ACTIVE,
				Muted:       track.LocalTrackPublication.IsMuted(),
				Remote:      false,
			},
		},
	}
	return res, nil
}
func (api *API) LocalTrackMute(req *pb_track.LocalTrackMuteReq) (*pb_track.LocalTrackMuteRes, error) {
	track := api.getLocalTrack(req.TrackHandle)
	track.LocalTrackPublication.SetMuted(req.Mute)
	res := &pb_track.LocalTrackMuteRes{}
	return res, nil
}
func (api *API) EnableRemoteTrack(req *pb_track.EnableRemoteTrackReq) (*pb_track.EnableRemoteTrackRes, error) {
	track := api.getRemoteTrack(req.TrackHandle)
	track.SetEnable(req.Enabled)
	res := &pb_track.EnableRemoteTrackRes{}
	return res, nil
}

func (api *API) EnableRemoteTrackPublication(req *pb_track.EnableRemoteTrackPublicationReq) (*pb_track.EnableRemoteTrackPublicationRes, error) {
	track := api.getRemoteTrack(req.RemoteTrackHandle)
	track.EnableTrackPubliciation(req.TrackPublicationSid, req.Enabled)
	res := &pb_track.EnableRemoteTrackPublicationRes{}
	return res, nil
}

func (api *API) UpdateRemoteTrackPublicationDimension(req *pb_track.UpdateRemoteTrackPublicationDimensionReq) (*pb_track.UpdateRemoteTrackPublicationDimensionRes, error) {
	track := api.getRemoteTrack(req.RemoteTrackHandle)
	track.UpdatePublicationDimension(req.TrackPublicationSid, req.Width, req.Height)
	res := &pb_track.UpdateRemoteTrackPublicationDimensionRes{}
	return res, nil
}

func (api *API) SetSubscribed(req *pb_track.SetSubscribedReq) (*pb_track.SetSubscribedRes, error) {
	track := api.getRemoteTrack(req.RemoteTrackHandle)
	track.SetSubscribed(req.PublicationSid, req.Subscribe)
	return nil, nil
}
