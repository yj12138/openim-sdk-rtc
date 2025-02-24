package sdk

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"log"
)

type Room struct {
	host        string
	token       string
	roomName    string
	owner       string
	livekitRoom *lksdk.Room

	listener OnRoomListener
	callBack *lksdk.RoomCallback
}

func (r *Room) ConnectByToken(host, token string) {
	r.host = host
	r.token = token
	livekitRoom, err := lksdk.ConnectToRoomWithToken(host, token, r.callBack, lksdk.WithAutoSubscribe(true))
	if err != nil {
		log.Println(err.Error())
		return
	}
	r.livekitRoom = livekitRoom
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

func (r *Room) checkConn() bool {
	if r.IsConnSuc() {
		return true
	} else {
		log.Panic("not connect to room")
		return false
	}
}

func (r *Room) GetRoomName() string {
	return r.roomName
}

func (r *Room) GetOwner() string {
	return r.owner
}

func (r *Room) GetLocalParticipant() *lksdk.LocalParticipant {
	return r.livekitRoom.LocalParticipant
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

func NewRoom(listener OnRoomListener) *Room {
	room := &Room{
		listener: listener,
		callBack: lksdk.NewRoomCallback(),
	}
	return room
}
