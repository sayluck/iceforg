package menu

import (
	"errors"
	"fmt"
	. "iceforg/app/log"
	"strings"

	"gopkg.in/go-playground/validator.v8"
)

func (m *MenuReq) CheckError(err error) []error {
	if err == nil {
		return nil
	}

	var errs []error
	for _, err := range err.(validator.ValidationErrors) {
		eStr := fmt.Sprintf("%s%s",
			err.Field,
			strings.Trim(err.ActualTag, ","))
		errs = append(errs, errors.New(eStr))
		Log.Errorf("user register validate[%s] error,parms[%s]", eStr, err.Param)
	}
	return errs
}
