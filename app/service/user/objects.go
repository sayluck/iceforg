package user

import "context"

// request
type UserRegister struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nickName"`
	context.Context
}

type UserLogin struct {
	UserID   string `json:"code"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	context.Context
}

// response
type UserDetail struct {
	UserID   string `json:"code"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
}
