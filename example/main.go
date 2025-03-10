package main

import (
	"flag"
	"log"

	"github.com/AllenDang/cimgui-go/backend"
	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
	"github.com/openimsdk/openim-rtc/example/ui"
)

var currentBackend backend.Backend[ebitenbackend.EbitenBackendFlags]

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	roomName := flag.String("r", "", "Room Name")
	particiantName := flag.String("p", "", "Participant Name")
	flag.Parse()

	if *roomName == "" || *particiantName == "" {
		flag.Usage()
		return
	}
	httpURL := "https://cloud-api.livekit.io/api/sandbox/connection-details"
	x_sandbox_id := "contextual-shell-22mdn5"
	ui.InitContext(httpURL, x_sandbox_id, *roomName, *particiantName)

	currentBackend, _ = backend.CreateBackend(ebitenbackend.NewEbitenBackend())
	currentBackend.SetAfterCreateContextHook(func() {
		implot.PlotCreateContext()
	})
	currentBackend.SetBeforeDestroyContextHook(func() {
		implot.PlotDestroyContext()
	})
	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	currentBackend.CreateWindow("OpenIM-Rtc-Demo", 1000, 800)
	ui.SetTitleCallBack(func(title string) {
		currentBackend.SetWindowTitle(title)
	})
	currentBackend.SetCloseCallback(func(b backend.Backend[ebitenbackend.EbitenBackendFlags]) {
		ui.Destory()
	})
	currentBackend.Run(func() {
		ui.Loop()
	})
}
