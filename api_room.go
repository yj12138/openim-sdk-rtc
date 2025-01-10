package sdk

import (
	"fmt"
	"sync"

	lksdk "github.com/livekit/server-sdk-go/v2"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

var roomMap sync.Map
var roomHandleInc int

func newRoomCallBack() *lksdk.RoomCallback {
	pb_roomCallBack := lksdk.NewRoomCallback()
	// TODO
	return pb_roomCallBack
}

func Connect(request *pb_room.ConnectRequest) *pb_room.ConnectResponse {
	asyncId := genAsyncId()
	go func() {
		room, err := lksdk.ConnectToRoomWithToken(request.Url, request.Token, newRoomCallBack(), lksdk.WithAutoSubscribe(request.Options.AutoSubscribe))
		if err != nil {
			EmitPanic(err.Error())
		}
		roomHandleInc++
		roomMap.Store(roomHandleInc, room)
		EmitConnectCallback()
	}()
	return &pb_room.ConnectResponse{
		AsyncId: asyncId,
	}
}

func Disconnect(request *pb_room.DisconnectRequest) *pb_room.DisconnectResponse {
	if val, ok := roomMap.Load(request.RoomHandle); ok {
		if room, ok := val.(*lksdk.Room); ok {
			room.Disconnect()
			EmitDisconnectCallback()
		} else {
			EmitPanic(fmt.Sprintf("not a room:%d", request.RoomHandle))
		}
	} else {
		EmitPanic(fmt.Sprintf("not find room handle:%d", request.RoomHandle))
	}
	return &pb_room.DisconnectResponse{
		AsyncId: genAsyncId(),
	}
}

func PublicTrack(request *pb_room.PublishTrackRequest) *pb_room.PublishTrackResponse {
    
	return nil
}

func UnpublishTrack(request *pb_room.UnpublishTrackRequest) *pb_room.UnpublishTrackResponse {
	return nil
}

func PublishData(request *pb_room.PublishDataRequest) *pb_room.PublishDataResponse {
	return nil
}

func SetSubscribed(request *pb_room.SetSubscribedRequest) *pb_room.SetSubscribedResponse {
	return nil
}

func UpdateLocalMetadata(request *pb_room.UpdateLocalMetadataRequest) *pb_room.UpdateLocalMetadataResponse {
	return nil
}

func UpdateLocalName(request *pb_room.UpdateLocalNameRequest) *pb_room.UpdateLocalNameResponse {
	return nil
}

func GetSessionStats(request *pb_room.GetSessionStatsRequest) *pb_room.GetSessionStatsResponse {
	return nil
}
