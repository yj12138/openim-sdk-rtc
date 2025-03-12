package ui

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/sdk"

	_io "github.com/openimsdk/openim-rtc/example/io"
)

var context *Context

type ConnectState string

var (
	ConnectNone    ConnectState = "Connect_None"
	Connecting     ConnectState = "Connecting"
	ConnectFailed  ConnectState = "Connect_Failed"
	ConnectSuccess ConnectState = "Connect_Success"
)

type CreateTokenResponse struct {
	ServerUrl        string `json:"serverUrl"`
	RoomName         string `json:"roomName"`
	ParticipantName  string `json:"participantName"`
	ParticipantToken string `json:"participantToken"`
}

type Context struct {
	httpURL             string
	x_Sandbox_ID        string
	serverUrl           string
	roomName            string
	participantIdentify string
	participantName     string
	participantToken    string
	ConnectState        ConnectState

	setWindowTitle func(string)

	Room             *sdk.Room
	LocalParticipant *sdk.LocalParticipant

	MicPhone *_io.MicPhone
	Speaker  *_io.Speaker
}

func (c *Context) connect() {
	c.ConnectState = Connecting
	data := map[string]string{
		"roomName": c.roomName,
		// "participantName": , 随机一个 identify
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Panic(jsonData)
		return
	}
	req, err := http.NewRequest("POST", c.httpURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Sandbox-ID", c.x_Sandbox_ID)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	var response CreateTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Panicln(err.Error())
		return
	}
	c.serverUrl = response.ServerUrl
	c.roomName = response.RoomName
	c.participantIdentify = response.ParticipantName
	c.participantToken = response.ParticipantToken
	log.Println(response.ServerUrl)
	log.Println(response.RoomName)
	log.Println(response.ParticipantName)
	log.Println(response.ParticipantToken)
	c.Room = sdk.ConnectByToken(c.serverUrl, c.participantToken, NewRoomListener())
	c.LocalParticipant = sdk.NewLocalParticipant(c.Room.LocalParticipant)
	c.setWindowTitle("Room:" + c.roomName)
	c.Room.LocalParticipant.SetName(c.participantName)
	go func() {
		time.Sleep(3 * time.Second)
		log.Println(c.Room.LocalParticipant.Name())
		log.Println(c.Room.LocalParticipant.Identity())
		log.Println(c.Room.LocalParticipant.SID())
		log.Println(c.Room.LocalParticipant.Attributes())
		log.Println(c.Room.LocalParticipant.IsMicrophoneEnabled())
		log.Println(c.Room.LocalParticipant.IsCameraEnabled())
		log.Println(c.Room.LocalParticipant.IsSpeaking())
		c.Room.LocalParticipant.SetName(c.participantName + " : new name )")
	}()

	c.ConnectState = ConnectSuccess
}

func (c *Context) PublishAudioTrack(track *sdk.LocalTrack) {
	_, err := c.LocalParticipant.PublishTrack(track, &lksdk.TrackPublicationOptions{
		Name:       track.Name,
		Source:     livekit.TrackSource_MICROPHONE,
		Encryption: livekit.Encryption_NONE,
	})
	if err != nil {
		log.Panic(err.Error())
	}
}

func (c *Context) SetWindowTitle(title string) {
	if c.setWindowTitle != nil {
		c.setWindowTitle(title)
	}
}

func InitContext(httpUrl string, x_Sandbox_ID string, roomName string, participantName string) {
	context = &Context{
		httpURL:         httpUrl,
		x_Sandbox_ID:    x_Sandbox_ID,
		roomName:        roomName,
		participantName: participantName,
		ConnectState:    ConnectNone,
		MicPhone:        _io.NewMicPhone(48000, 1),
		Speaker:         _io.NewSpeaker(48000, 1),
	}
	go context.connect()
}

func SetTitleCallBack(callBack func(title string)) {
	context.setWindowTitle = callBack
}
