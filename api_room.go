package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/proto/go/room"
)

func Connect(request *room.ConnectRequest) *room.ConnectResponse {
	_, err := lksdk.ConnectToRoomWithToken(request.Url, request.Token, NewRoomCallBack(), lksdk.WithAutoSubscribe(request.Options.AutoSubscribe))
	if err != nil {
		return nil
	}
	return &room.ConnectResponse{
		AsyncId: genAsyncId(),
	}
}

func Disconnect(request *room.DisconnectRequest) *room.DisconnectResponse {
	return nil
}

func PublicTrack(request *room.PublishTrackRequest) *room.PublishTrackResponse {
	return nil
}

func UnpublishTrack(request *room.UnpublishTrackRequest) *room.UnpublishTrackResponse {
	return nil
}

func PublishData(request *room.PublishDataRequest) *room.PublishDataResponse {
	return nil
}

func SetSubscribed(request *room.SetSubscribedRequest) *room.SetSubscribedResponse {
	return nil
}

func UpdateLocalMetadata(request *room.UpdateLocalMetadataRequest) *room.UpdateLocalMetadataResponse {
	return nil
}

func UpdateLocalName(request *room.UpdateLocalNameRequest) *room.UpdateLocalNameResponse {
	return nil
}

func GetSessionStats(request *room.GetSessionStatsRequest) *room.GetSessionStatsResponse {
	return nil
}
