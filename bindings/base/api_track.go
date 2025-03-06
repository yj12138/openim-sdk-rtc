package base

import (
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	pb_track_publication "github.com/openimsdk/openim-rtc/proto/go/track_publication"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) CreateVideoTrack(req *pb_track.CreateVideoTrackRequest) (*pb_track.CreateVideoTrackResponse, error) {
	track := sdk.NewVideoTrack()
	trackHandle := api.storeObj(track)
	res := &pb_track.CreateVideoTrackResponse{
		Track: &pb_track.OwnedTrack{
			Handle: &pb_handle.FfiOwnedHandle{Id: trackHandle},
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
func (api *API) CreateAudioTrack(req *pb_track.CreateAudioTrackRequest) (*pb_track.CreateAudioTrackResponse, error) {
	track := sdk.NewAudioTrack(req.Name)
	trackHandle := api.storeObj(track)
	res := &pb_track.CreateAudioTrackResponse{
		Track: &pb_track.OwnedTrack{
			Handle: &pb_handle.FfiOwnedHandle{Id: trackHandle},
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
func (api *API) LocalTrackMute(req *pb_track.LocalTrackMuteRequest) (*pb_track.LocalTrackMuteResponse, error) {
	track := api.getLocalTrack(req.TrackHandle)
	track.LocalTrackPublication.SetMuted(req.Mute)
	res := &pb_track.LocalTrackMuteResponse{}
	return res, nil
}
func (api *API) EnableRemoteTrack(req *pb_track.EnableRemoteTrackRequest) (*pb_track.EnableRemoteTrackResponse, error) {
	track := api.getRemoteTrack(req.TrackHandle)
	track.SetEnable(req.Enabled)
	res := &pb_track.EnableRemoteTrackResponse{}
	return res, nil
}

func (api *API) EnableRemoteTrackPublication(req *pb_track_publication.EnableRemoteTrackPublicationRequest) (*pb_track_publication.EnableRemoteTrackPublicationResponse, error) {
	// track := api.GetTrackPublished(req.)
	// track.EnableTrackPubliciation(req.TrackPublicationSid, req.Enabled)
	res := &pb_track_publication.EnableRemoteTrackPublicationResponse{}
	return res, nil
}

func (api *API) UpdateRemoteTrackPublicationDimension(req *pb_track_publication.UpdateRemoteTrackPublicationDimensionRequest) (*pb_track_publication.UpdateRemoteTrackPublicationDimensionResponse, error) {
	// track := api.getRemoteTrack(req.RemoteTrackHandle)
	// track.UpdatePublicationDimension(req.TrackPublicationSid, req.Width, req.Height)
	res := &pb_track_publication.UpdateRemoteTrackPublicationDimensionResponse{}
	return res, nil
}

func (api *API) SetSubscribed(req *pb_room.SetSubscribedRequest) (*pb_room.SetSubscribedResponse, error) {
	// track := api.get(req.PublicationHandle)
	// track.SetSubscribed(req.PublicationSid, req.Subscribe)
	return nil, nil
}
