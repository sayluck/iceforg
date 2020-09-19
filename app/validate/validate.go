package validate

import (
	"context"
	"errors"
	"fmt"
	"iceforg/app/log"
	"strings"
	"sync"

	"gopkg.in/go-playground/validator.v8"
)

var (
	v    *validator.Validate
	once sync.Once
)

type Validater interface {
	SetContext(context.Context)
	CustomValidate() error
}

func InitValidate() {
	once.Do(func() {
		v = validator.New(&validator.Config{
			TagName:      "validate",
			FieldNameTag: "json",
		})
	})

}

func GetInstance() *validator.Validate {
	if v == nil {
		InitValidate()
	}
	return v
}

func ValidateStruct(ctx context.Context, val Validater) []error {
	var retErrs []error
	val.SetContext(ctx)
	errs := analysisError(ctx, GetInstance().Struct(val))
	err := val.CustomValidate()
	if errs != nil {
		retErrs = append(retErrs, errs...)
	}
	if err != nil {
		retErrs = append(retErrs, err)
	}
	return retErrs
}

func analysisError(ctx context.Context, err error) []error {
	if err == nil {
		return nil
	}
	var errs []error
	for _, err := range err.(validator.ValidationErrors) {
		eStr := fmt.Sprintf("field [%s] is %s",
			err.Field,
			strings.Trim(err.ActualTag, ","))
		errs = append(errs, errors.New(eStr))
		log.IceLog.Errorf(ctx, "validate[%s] error,parms[%s]", eStr, err.Param)
	}
	return errs
}
