package storage

import (
	log "github.com/ctlog"
	"github.com/iceforg/warehouse/libs/utils"
	"os"
	"path"
)

func GenerateStoragePath(bcType, name, version string) string {
	switch bcType {
	case utils.FABRIC:
		{
			p := generatePathByOptions(name, version)
			if p == "" {
				return ""
			}
			return path.Join(utils.CONTRACTFATHERDIR, bcType, p)
		}
	default:
		return ""
	}
}

func generatePathByOptions(options ...string) string {
	var p string
	for _, str := range options {
		if str == "" {
			return ""
		}
		p = path.Join(p, utils.StringEncryption(str))
	}
	return p
}

func CreateContractDirector(path string) error {
	// only read
	log.Infof("create contract dir:%s\n", path)
	return os.MkdirAll(path, 665)
}

func removeContractDirector(path string) error {
	log.Infof("remove contract dir:%s\n", path)
	return os.Remove(path)
}
