package common

import "fmt"

func (m *BaseReq) CheckCreator() error {
	if m.Creator == "" {
		return fmt.Errorf("field [creator] is required")
	}
	return nil
}

func (m *BaseReq) CheckModifier() error {
	if m.Modifier == "" {
		return fmt.Errorf("field [modifier] is required")
	}
	return nil
}
