package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AllenDang/cimgui-go/backend"
	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/openimsdk/openim-rtc/example/ui"

	"github.com/openimsdk/openim-rtc/example/core"
)

var currentBackend backend.Backend[ebitenbackend.EbitenBackendFlags]

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	host := flag.String("host", "http://127.0.0.1:7880", "livekit host url")
	apiKey := flag.String("key", "APImMfGZ6ECRXbd", "livekit api key")
	apiSecret := flag.String("secret", "3Uxe8kZorbAz59UDKQe23XjM5alE9T0OyCnSNXjQ9r9", "livekit api secret")
	roomName := flag.String("r", "", "Room Name")
	idenfify := flag.String("i", "", "Identify")
	flag.Parse()

	if *roomName == "" || *idenfify == "" {
		flag.Usage()
		return
	}

	core.Init(*host, *apiKey, *apiSecret)
	core.ConnectRoom(*roomName, *idenfify)

	currentBackend, _ = backend.CreateBackend(ebitenbackend.NewEbitenBackend())
	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	currentBackend.CreateWindow(fmt.Sprintf("RoomName:%s - Identify:%s", *roomName, *idenfify), 800, 500)
	currentBackend.SetCloseCallback(func(b backend.Backend[ebitenbackend.EbitenBackendFlags]) {
		ui.Destory()
	})
	currentBackend.Run(ui.GUILoop)
}
