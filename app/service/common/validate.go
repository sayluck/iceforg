package common

import (
	"iceforg/app/validate"
)

func (m *BaseReq) CheckCreator() error {
	return validate.GetInstance().Field(m.Creator, "required")
}
