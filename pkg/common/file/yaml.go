package file

import (
	yamlv2 "gopkg.in/yaml.v2"
	"io/ioutil"
)

type yaml struct {
	filePath string
}

func (y *yaml) LoadFile(val interface{}) error {
	var (
		err  error
		data []byte
	)

	data, err = ioutil.ReadFile(y.filePath)
	if err != nil {
		return err
	}

	return yamlv2.Unmarshal(data, val)
}
