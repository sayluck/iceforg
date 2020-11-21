package menu

import (
	"context"
	"fmt"
)

func (m *MenuAddReq) SetContext(ctx context.Context) {
	m.Context = ctx
}

func (u *MenuAddReq) CustomValidate() error {
	err := u.checkSupCode()
	return err
}

func (u *MenuAddReq) checkSupCode() error {
	if u.Level > 1 && u.SupCode == "" {
		return fmt.Errorf("field [supCode] is required,when level > 1")
	}
	return nil
}
