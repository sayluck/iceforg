package user

import (
	"context"
	"iceforg/app/log"
	appRPC "iceforg/app/rpc"
	"iceforg/modules/userCenter"
	"iceforg/pkg/registry/rpc"
)

var userCli rpc.RpcClient

func init() {
	userCli = appRPC.RpcClient.CreateClient("user")
}

func (u *UserRegister) Register() (string, error) {
	registerReq := &userCenter.UserRegister{
		UserName: u.UserName,
		Password: u.Password,
		NickName: u.NickName,
	}
	var (
		code string
		err  error
	)
	err = userCli.Call(u.Context, "Register", registerReq, &code)
	if err != nil {
		log.IceLog.Errorf(u.Context, "register user(name=%s) failed", u.UserName)
	}
	return code, err
}

func Detail(ctx context.Context, code string) (*UserDetail, error) {
	var (
		err    error
		detail *UserDetail
	)
	err = userCli.Call(ctx, "Detail", &code, &detail)
	if err != nil {
		log.IceLog.Errorf(ctx, "detail user(code=%s) failed", code)
	}
	return detail, err
}

func (u *UserLogin) Login() (string, error) {
	var (
		err   error
		token string
	)

	userLogin := &userCenter.UserLogin{
		UserName: u.UserName,
		Password: u.Password,
	}

	err = userCli.Call(u.Context, "Login", userLogin, &token)
	if err != nil {
		log.IceLog.Errorf(u.Context, "longin with user(name=%s) failed", u.UserName)
	}
	return token, err
}
