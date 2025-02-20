package base

import (
	"github.com/openimsdk/openim-rtc/proto/go/event"
)

var funcMap = map[event.FuncRequestEventName]callFunc{}
