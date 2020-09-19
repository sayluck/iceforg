package user

import (
	"context"
)

func (u *UserRegister) SetContext(ctx context.Context) {
	u.Context = ctx
}

func (u *UserRegister) CustomValidate() error {
	return nil
}
