package menu

import (
	"context"
)

func (m *MenuReq) SetContext(ctx context.Context) {
	m.Context = ctx
}

func (u *MenuReq) CustomValidate() error {
	return u.BaseReq.CheckCreator()
}
