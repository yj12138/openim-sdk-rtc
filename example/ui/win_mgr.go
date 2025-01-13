package ui

import "github.com/AllenDang/cimgui-go/imgui"

var windowMap = make(map[string]IWindow)

type IWindow interface {
	GetId() string
	Start()
	Update()
	Destroy()
}

type WindowBase struct {
	Id    string
	Pos   imgui.Vec2 // pos
	Size  imgui.Vec2 // width,height
	Title string     // title name
	Open  bool
}

func (w *WindowBase) GetId() string {
	return w.Id
}

func (w *WindowBase) Start() {

}

func (w *WindowBase) Update() {

}

func (w *WindowBase) Destroy() {

}

func showWindow(win IWindow) {
	win.Start()
	windowMap[win.GetId()] = win
}

func DrawAllWindow() {
	for _, win := range windowMap {
		win.Update()
	}
}
