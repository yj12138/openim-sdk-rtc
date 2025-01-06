package main

import (
	"log"

	"github.com/AllenDang/cimgui-go/backend"
	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/AllenDang/cimgui-go/imgui"
)

var currentBackend backend.Backend[ebitenbackend.EbitenBackendFlags]

func init() {
	log.SetFlags(log.Llongfile)
}

func loop() {
	// ui.DrawMainMenu()
	// ui.DrawAllWindow()
}

func main() {
	// userId := flag.String("userId", "", "User ID ")
	// token := flag.String("token", "", "User Token")
	// flag.Parse()

	// if *userId == "" || *token == "" {
	// 	flag.Usage()
	// 	return
	// }

	// nickName := ui.Login(*userId, *token)

	currentBackend, _ = backend.CreateBackend(ebitenbackend.NewEbitenBackend())
	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	currentBackend.CreateWindow("OpenIM-RTC-Demo", 1200, 900)
	currentBackend.SetCloseCallback(func(b backend.Backend[ebitenbackend.EbitenBackendFlags]) {
	})

	currentBackend.Run(loop)
}
