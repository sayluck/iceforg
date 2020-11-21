package api

import (
	"fmt"
)

const (
	succCode = "2000"

	SystemErr   = "3001"
	SystemPanic = "3002"

	ParamsErr    = "4001"
	OperationErr = "4002"

	UserInvalidToken = "5001"
)

type Resp struct {
	ReqID  string      `json:"reqID"`
	Data   interface{} `json:"data"`
	RetMsg retMsg      `json:"retMsg"`
	Status string      `json:"status"`
}

type retMsg struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func RespFailed(code string, msg ...string) *Resp {
	return &Resp{
		ReqID:  "",
		Data:   nil,
		Status: "error",
		RetMsg: retMsg{
			Code: code,
			Msg:  fmt.Sprintf("%v", msg),
		},
	}
}

func RespSucc(data interface{}, msg ...string) *Resp {
	if len(msg) <= 0 {
		msg = []string{"succ"}
	}
	return &Resp{
		ReqID:  "",
		Data:   data,
		Status: "ok",
		RetMsg: retMsg{
			Code: succCode,
			Msg:  fmt.Sprintf("%v", msg),
		},
	}
}
