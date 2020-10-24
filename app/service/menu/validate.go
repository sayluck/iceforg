package menu

import (
	"context"
)

func (m *MenuAddReq) SetContext(ctx context.Context) {
	m.Context = ctx
}

func (u *MenuAddReq) CustomValidate() error {
	return u.BaseReq.CheckCreator()
}
