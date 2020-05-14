package validate

import (
	"sync"

	"gopkg.in/go-playground/validator.v8"
)

var (
	v    *validator.Validate
	once sync.Once
)

type Validater interface {
	CheckError(error) []error
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

type CheckErr func(err error) error

func ValidateStruct(val Validater) []error {
	err := GetInstance().Struct(val)
	return val.CheckError(err)
}
