package common

import (
	"iceforg/pkg/common/file"
)

func LoadFile(filePath string, v interface{}) error {
	return file.LoadFileFactory(filePath).LoadFile(v)
}
