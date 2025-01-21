package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"log"
	// pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
)

type Room struct {
	host        string
	token       string
	roomName    string
	owner       string
	livekitRoom *lksdk.Room

	callBack *lksdk.RoomCallback
}

func (r *Room) ConnectByToken(host, token, roomName, identify string) {
	r.host = host
	r.token = token
	go func() {
		livekitRoom, err := lksdk.ConnectToRoomWithToken(host, token, r.callBack, lksdk.WithAutoSubscribe(true))
		if err != nil {
			log.Println(err.Error())
			return
		}
		r.livekitRoom = livekitRoom
	}()
}

func (r *Room) ConnectBySecret(host, apiKey, apiSecret, roomName, identify string) {
	r.roomName = roomName
	r.owner = identify
	room, err := lksdk.ConnectToRoom(host, lksdk.ConnectInfo{
		APIKey:              apiKey,
		APISecret:           apiSecret,
		RoomName:            roomName,
		ParticipantIdentity: identify,
	}, r.callBack, lksdk.WithAutoSubscribe(true))
	if err != nil {
		panic(err)
	}
	r.livekitRoom = room
}

func (r *Room) Disconnect() {
	if r.livekitRoom != nil {
		r.livekitRoom.Disconnect()
	}
	r.livekitRoom = nil
}

func (r *Room) IsConnSuc() bool {
	return r.livekitRoom != nil
}

func (r *Room) PublicTrack(track *Track) {
	if track.liveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	if r.checkConn() {
		log.Println("public track-------------", track.name, track.mimeType)
		_, err := r.livekitRoom.LocalParticipant.PublishTrack(track.liveKitTrack, &lksdk.TrackPublicationOptions{
			VideoWidth:  track.videoWidth,
			VideoHeight: track.videoHeight,
			Name:        track.name,
		})
		if err != nil {
			log.Panic(err)
		}
	}
}

func (r *Room) UnpublishTrack(track *Track) {
	if track.liveKitTrack == nil {
		log.Panic("track is nil")
		return
	}
	if r.checkConn() {
		err := r.livekitRoom.LocalParticipant.UnpublishTrack(track.liveKitTrack.StreamID())
		if err != nil {
			log.Panic(err)
		}
	}
}

func (r *Room) checkConn() bool {
	if r.IsConnSuc() {
		return true
	} else {
		log.Panic("not connect to room")
		return false
	}
}

func (r *Room) PublishData(data string) {
	if r.checkConn() {
		r.livekitRoom.LocalParticipant.PublishDataPacket(lksdk.UserData([]byte(data)), lksdk.WithDataPublishReliable(true))
	}
}

func (r *Room) SetSubscribed() {
}

func (r *Room) UpdateLocalMetadata() {
}

func (r *Room) UpdateLocalName() {
}

func (r *Room) GetSessionStats() {
}

func (r *Room) GetRoomName() string {
	return r.roomName
}

func (r *Room) GetOwner() string {
	return r.owner
}

func (r *Room) GetAllParticipantId() []string {
	res := make([]string, 0)
	if r.checkConn() {
		rps := r.livekitRoom.GetRemoteParticipants()
		for _, rp := range rps {
			res = append(res, rp.Identity())
		}
	}
	return res
}

func NewRoom(callBack *lksdk.RoomCallback) *Room {
	room := &Room{
		callBack: callBack,
	}
	return room
}
