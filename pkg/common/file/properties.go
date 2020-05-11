package file

import (
	"reflect"
	"strings"

	plib "github.com/magiconair/properties"
)

type properties struct {
	filePath string
}

func (p *properties) LoadFile(val interface{}) error {
	pp, err := plib.LoadFile(p.filePath, plib.UTF8)
	if err != nil {
		return err
	}

	message_value := reflect.ValueOf(val).Elem().Elem()

	for i := 0; i < message_value.NumField(); i++ {
		fieldInfo := message_value.Type().Field(i)
		name := strings.ToLower(fieldInfo.Name)

		if name == PROPERTIES {
			message_value.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(pp))
		}
	}
	return nil
}
