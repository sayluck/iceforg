package api

import (
	"fmt"
)

const (
	succCode    = "2000"

	SystemErr   = "3001" 
	SystemPanic = "3002"

	ParamsErr   = "4001"
	OperationErr   = "4002"
)

type Resp struct {
	ReqID  string      `json:"reqID"`
	Data   interface{} `json:"data"`
	RetMsg retMsg      `json:"retMsg"`
	IsSucc bool        `json:"isSucc"`
}

type retMsg struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func RespFailed(code string, msg ...string) *Resp {
	return &Resp{
		ReqID:  "",
		Data:   nil,
		IsSucc: false,
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
		IsSucc: true,
		RetMsg: retMsg{
			Code: succCode,
			Msg:  fmt.Sprintf("%v", msg),
		},
	}
}
