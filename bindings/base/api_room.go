package base

import (
	// lksdk "github.com/livekit/server-sdk-go/v2"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) Connect(req *pb_room.ConnectRequest) (*pb_room.ConnectResponse, error) {
	listener := NewRoomListener()
	r := sdk.NewRoom(listener)
	roomHandle := api.storeObj(r)
	listener.RoomHandle = roomHandle
	r.ConnectByToken(req.Url, req.Token)
	// participant := sdk.NewLocalParticipant(r.GetLocalParticipant())
	// localParticipantHandle := api.storeObj(participant)
	res := &pb_room.ConnectResponse{
		AsyncId: 0,
	}
	return res, nil
}

func (api *API) Disconnect(req *pb_room.DisconnectRequest) (*pb_room.DisconnectResponse, error) {
	room := api.getRoom(req.RoomHandle)
	room.Disconnect()
	return &pb_room.DisconnectResponse{}, nil
}

func (api *API) GetSessionStats(req *pb_room.GetSessionStatsRequest) (*pb_room.GetSessionStatsResponse, error) {
	// room := api.getRoom(req.RoomHandle)
	res := &pb_room.GetSessionStatsResponse{
		// State: room.GetConnectState(),
	}
	return res, nil
}
