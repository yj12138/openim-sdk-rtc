package ui

import "github.com/openimsdk/openim-rtc/sdk"

func init() {
	sdk.SetRoomListener(NewRoomListener())
}