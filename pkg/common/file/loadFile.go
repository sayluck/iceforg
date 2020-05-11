package file

import (
	"path"
	"strings"
)

const (
	// file type
	YAML       = "yaml"
	TOML       = "toml"
	PROPERTIES = "properties"
)

type fileTools interface {
	// loadfile
	LoadFile(val interface{}) error
}

func LoadFileFactory(fielPath string) fileTools {
	// judge by file type
	switch strings.Trim(path.Ext(fielPath), ".") {
	case TOML:
		return &toml{
			filePath: fielPath,
		}
	case YAML:
		return &yaml{
			filePath: fielPath,
		}
	case PROPERTIES:
		return &properties{
			filePath: fielPath,
		}
	default:
		return &yaml{
			filePath: fielPath,
		}
	}
}
