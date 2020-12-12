package user

import (
	"context"
	"iceforg/modules/userCenter"
)

func (u *userCenter.UserRegister) SetContext(ctx context.Context) {
	u.Context = ctx
}

func (u *userCenter.UserRegister) CustomValidate() error {
	return nil
}
