package base

import (
	"github.com/openimsdk/openim-rtc/proto/go/error"
)

var (
	Err_None            = newError(error.ErrCode_Err_None)
	Err_Proto_Marshal   = newError(error.ErrCode_Err_Proto_Marshal)
	Err_Proto_UnMarshal = newError(error.ErrCode_Err_Proto_Unmarshal)
	Err_SDK_Internal    = newError(error.ErrCode_Err_Sdk_Internal)
)

type Error struct {
	Code   error.ErrCode
	ErrMsg string
}

func (e *Error) New(errMsg string) Error {
	return Error{
		Code:   e.Code,
		ErrMsg: errMsg,
	}
}

func newError(code error.ErrCode) Error {
	return Error{
		Code:   code,
		ErrMsg: "",
	}
}
