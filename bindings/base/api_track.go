package base

import (
	"log"

	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	pb_track_publication "github.com/openimsdk/openim-rtc/proto/go/track_publication"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) CreateVideoTrack(req *pb_track.CreateVideoTrackRequest) *pb_track.CreateVideoTrackResponse {
	videoSource := api.getVideoSource(req.SourceHandle)
	track := sdk.NewVideoTrack(req.Name, videoSource)
	res := &pb_track.CreateVideoTrackResponse{
		Track: &pb_track.OwnedTrack{
			Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(track)},
			Info: &pb_track.TrackInfo{
				// Sid:         track.LocalTrackPublication.SID(),
				// Name:        track.LocalTrackPublication.Name(),
				// Kind:        sdk.ConvertTrackKind(track.LocalTrackPublication.Kind()),
				// StreamState: pb_track.StreamState_STATE_ACTIVE,
				// Muted:       track.LocalTrackPublication.IsMuted(),
				Remote: false,
			},
		},
	}
	return res
}
func (api *API) CreateAudioTrack(req *pb_track.CreateAudioTrackRequest) *pb_track.CreateAudioTrackResponse {
	source := api.getAudioSource(req.SourceHandle)
	track := sdk.NewAudioTrack(req.Name, source)
	trackHandle := api.storeObj(track)
	res := &pb_track.CreateAudioTrackResponse{
		Track: &pb_track.OwnedTrack{
			Handle: &pb_handle.FfiOwnedHandle{Id: trackHandle},
			Info: &pb_track.TrackInfo{
				// Sid:         track.StreamID(),
				// Name:        track.Name(),
				// Kind:        sdk.ConvertTrackKind(track.LocalTrackPublication.Kind()),
				// StreamState: pb_track.StreamState_STATE_ACTIVE,
				// Muted:       track.LocalTrackPublication.IsMuted(),
				Remote: false,
			},
		},
	}
	return res
}
func (api *API) LocalTrackMute(req *pb_track.LocalTrackMuteRequest) *pb_track.LocalTrackMuteResponse {
	track := api.getLocalTrack(req.TrackHandle)
	track.Publication.SetMuted(req.Mute)
	res := &pb_track.LocalTrackMuteResponse{
		Muted: track.Publication.IsMuted(),
	}
	return res
}
func (api *API) EnableRemoteTrack(req *pb_track.EnableRemoteTrackRequest) *pb_track.EnableRemoteTrackResponse {
	pub := api.getRemoteTackPublication(req.TrackHandle)
	pub.SetEnabled(req.Enabled)
	res := &pb_track.EnableRemoteTrackResponse{
		Enabled: pub.IsEnabled(),
	}
	return res
}

func (api *API) EnableRemoteTrackPublication(req *pb_track_publication.EnableRemoteTrackPublicationRequest) *pb_track_publication.EnableRemoteTrackPublicationResponse {

	res := &pb_track_publication.EnableRemoteTrackPublicationResponse{}
	return res
}

func (api *API) UpdateRemoteTrackPublicationDimension(req *pb_track_publication.UpdateRemoteTrackPublicationDimensionRequest) *pb_track_publication.UpdateRemoteTrackPublicationDimensionResponse {
	pub := api.getRemoteTackPublication(req.TrackPublicationHandle)
	pub.SetVideoDimensions(req.Width, req.Height)
	res := &pb_track_publication.UpdateRemoteTrackPublicationDimensionResponse{}
	return res
}

func (api *API) SetSubscribed(req *pb_room.SetSubscribedRequest) *pb_room.SetSubscribedResponse {
	track := api.getRemoteTackPublication(req.PublicationHandle)
	err := track.SetSubscribed(req.Subscribe)
	if err != nil {
		log.Println("SetSubscribed error:", err)
	}
	return &pb_room.SetSubscribedResponse{}
}

func (api *API) GetStats(req *pb_track.GetStatsRequest) *pb_track.GetStatsResponse {
	asyncId := api.nextAsyncId()
	go func() {
	}()
	return &pb_track.GetStatsResponse{
		AsyncId: asyncId,
	}
}
