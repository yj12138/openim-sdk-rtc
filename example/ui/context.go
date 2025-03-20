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
	_io "github.com/openimsdk/openim-rtc/example/io"
	pb_audio_frame "github.com/openimsdk/openim-rtc/proto/go/audio_frame"
	"github.com/openimsdk/openim-rtc/sdk"
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

	MicPhone    *_io.MicPhone
	Speaker     *_io.Speaker
	audioSource *sdk.AudioSource
}

func (c *Context) connect() {
	c.ConnectState = Connecting
	data := map[string]string{
		"roomName":        c.roomName,
		"participantName": "test-go-client", // particiant identify
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
	go func() {
		time.Sleep(1 * time.Second)
		log.Println(c.Room.LocalParticipant.Name())
		log.Println(c.Room.LocalParticipant.Identity())
		log.Println(c.Room.LocalParticipant.SID())
		log.Println(c.Room.LocalParticipant.Attributes())
		log.Println(c.Room.LocalParticipant.IsMicrophoneEnabled())
		log.Println(c.Room.LocalParticipant.IsCameraEnabled())
		log.Println(c.Room.LocalParticipant.IsSpeaking())
		c.Room.LocalParticipant.SetName(c.participantName)
	}()
	c.ConnectState = ConnectSuccess

	context.Speaker.Start()
	context.MicPhone.CallBack = (func(data []byte, frameCount uint32) {
		appendRawAudioFrame(data, frameCount)
		// TODO echo cancle
		// appendClearAudioFrame(frame, frameCount)
		if context.audioSource != nil {
			samplesPreChannel := len(data) / 2 * int(c.MicPhone.Channels)
			context.audioSource.CaptureFrame(data, c.MicPhone.Channels, c.MicPhone.SampleRate, uint32(samplesPreChannel))
			if err != nil {
				log.Println(err.Error())
			}
		}
	})
}

func (c *Context) SendData(data string) {
	go func() {
		err := context.LocalParticipant.SendData("Test", []byte("golang hello"), true, []string{})
		if err != nil {
			log.Println(err.Error())
		}
	}()
}

func (c *Context) PublishAudioTrack() {
	c.audioSource = sdk.NewAudioSource(pb_audio_frame.AudioSourceType_AUDIO_SOURCE_NATIVE, context.MicPhone.SampleRate, context.MicPhone.Channels)
	track := sdk.NewAudioTrack("micphone audio track", c.audioSource)
	_, err := c.LocalParticipant.PublishTrack(track, &lksdk.TrackPublicationOptions{
		Name:       track.Name,
		Source:     livekit.TrackSource_MICROPHONE,
		Encryption: livekit.Encryption_NONE,
	})
	if err != nil {
		log.Panic(err.Error())
	}
	context.MicPhone.Start()
}

func (c *Context) SetWindowTitle(title string) {
	if c.setWindowTitle != nil {
		c.setWindowTitle(title)
	}
}

func InitContext(httpUrl string, x_Sandbox_ID string, roomName string, participantName string) {
	sampleRate := 48000
	context = &Context{
		httpURL:         httpUrl,
		x_Sandbox_ID:    x_Sandbox_ID,
		roomName:        roomName,
		participantName: participantName,
		ConnectState:    ConnectNone,
		MicPhone:        _io.NewMicPhone(uint32(sampleRate), 1),
		Speaker:         _io.NewSpeaker(uint32(sampleRate), 1),
	}
	go context.connect()
}

func SetTitleCallBack(callBack func(title string)) {
	context.setWindowTitle = callBack
}
