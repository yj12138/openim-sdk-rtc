package base

import (
	// lksdk "github.com/livekit/server-sdk-go/v2"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) Connect(req *pb_room.ConnectReq) (*pb_room.ConnectRes, error) {
	listener := NewRoomListener()
	r := sdk.NewRoom(listener)
	roomHandle := api.storeObj(r)
	listener.RoomHandle = roomHandle
	r.ConnectByToken(req.Url, req.Token)
	participant := sdk.NewLocalParticipant(r.GetLocalParticipant())
	localParticipantHandle := api.storeObj(participant)
	res := &pb_room.ConnectRes{
		RoomHandle:             roomHandle,
		LocalParticipantHandle: localParticipantHandle,
	}
	return res, nil
}

func (api *API) Disconnect(req *pb_room.DisconnectReq) (*pb_room.DisconnectRes, error) {
	room := api.getRoom(req.RoomHandle)
	room.Disconnect()
	return &pb_room.DisconnectRes{}, nil
}

func (api *API) GetSessionStats(req *pb_room.GetSessionStatsReq) (*pb_room.GetSessionStatsRes, error) {

	return nil, nil
}
