package base

import (
	"github.com/openimsdk/openim-rtc/proto/go/e2ee"
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) Connect(req *pb_room.ConnectRequest) *pb_room.ConnectResponse {
	asyncId := api.nextAsyncId()
	go func() {
		listener := NewRoomListener()
		room := sdk.ConnectByToken(req.Url, req.Token, listener)
		roomHandle := api.storeObj(room)
		listener.RoomHandle = roomHandle
		localParticipant := sdk.NewLocalParticipant(room.LocalParticipant)
		localParticipantHandle := api.storeObj(localParticipant)
		rps := room.GetRemoteParticipants()
		participantWithTracks := make([]*pb_room.ConnectCallback_ParticipantWithTracks, 0)
		for _, rp := range rps {
			participant := &pb_participant.OwnedParticipant{
				Info: &pb_participant.ParticipantInfo{
					Sid:        rp.SID(),
					Name:       rp.Name(),
					Identity:   rp.Identity(),
					Metadata:   rp.Metadata(),
					Attributes: rp.Attributes(),
					Kind:       pb_participant.ParticipantKind(rp.Kind()),
				},
			}
			participant.Handle = &pb_handle.FfiOwnedHandle{Id: api.storeObj(rp)}
			publications := rp.TrackPublications()
			pubs := make([]*pb_track.OwnedTrackPublication, 0)
			for _, pub := range publications {
				pubs = append(pubs, &pb_track.OwnedTrackPublication{
					Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(pub)},
					Info: &pb_track.TrackPublicationInfo{
						Sid:            pub.SID(),
						Name:           pub.Name(),
						Kind:           sdk.ConvertTrackKind(pub.Kind()),
						Source:         pb_track.TrackSource(pub.Source()),
						Width:          pub.TrackInfo().Width,
						Height:         pub.TrackInfo().Height,
						MimeType:       pub.MimeType(),
						Muted:          pub.IsMuted(),
						Remote:         true,
						EncryptionType: e2ee.EncryptionType_NONE,
					},
				})
			}
			participantWithTracks = append(participantWithTracks, &pb_room.ConnectCallback_ParticipantWithTracks{
				Participant:  participant,
				Publications: pubs,
			})
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_Connect{
				Connect: &pb_room.ConnectCallback{
					AsyncId: asyncId,
					Message: &pb_room.ConnectCallback_Result_{
						Result: &pb_room.ConnectCallback_Result{
							Room: &pb_room.OwnedRoom{
								Handle: &pb_handle.FfiOwnedHandle{
									Id: roomHandle,
								},
								Info: &pb_room.RoomInfo{},
							},
							LocalParticipant: &pb_participant.OwnedParticipant{
								Handle: &pb_handle.FfiOwnedHandle{Id: localParticipantHandle},
								Info: &pb_participant.ParticipantInfo{
									Sid:        room.LocalParticipant.SID(),
									Name:       room.LocalParticipant.Name(),
									Identity:   room.LocalParticipant.Identity(),
									Metadata:   room.LocalParticipant.Metadata(),
									Attributes: room.LocalParticipant.Attributes(),
									Kind:       pb_participant.ParticipantKind(room.LocalParticipant.Kind()),
								},
							},
							Participants: participantWithTracks,
						},
					},
				},
			},
		})
	}()
	res := &pb_room.ConnectResponse{
		AsyncId: asyncId,
	}
	return res
}

func (api *API) Disconnect(req *pb_room.DisconnectRequest) *pb_room.DisconnectResponse {
	asyncId := api.nextAsyncId()
	room := api.getRoom(req.RoomHandle)
	go func() {
		room.Disconnect()
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_Disconnect{
				Disconnect: &pb_room.DisconnectCallback{
					AsyncId: asyncId,
				},
			},
		})
	}()
	return &pb_room.DisconnectResponse{
		AsyncId: asyncId,
	}
}

func (api *API) GetSessionStats(req *pb_room.GetSessionStatsRequest) *pb_room.GetSessionStatsResponse {
	asyncId := api.nextAsyncId()
	room := api.getRoom(req.RoomHandle)
	go func() {
		// TODO
		room.GetConnectState()
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_GetSessionStats{
				GetSessionStats: &pb_room.GetSessionStatsCallback{
					AsyncId: asyncId,
					Message: &pb_room.GetSessionStatsCallback_Result_{
						Result: &pb_room.GetSessionStatsCallback_Result{},
					},
				},
			},
		})
	}()
	res := &pb_room.GetSessionStatsResponse{
		AsyncId: asyncId,
	}
	return res
}

func (api *API) Dispose(req *pb_ffi.DisposeRequest) *pb_ffi.DisposeResponse {
	asyncId := api.nextAsyncId()
	go func() {
		// TODO
	}()
	return &pb_ffi.DisposeResponse{
		AsyncId: asyncId,
	}
}
