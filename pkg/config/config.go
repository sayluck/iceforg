package config

import (
	"fmt"
	"iceforg/pkg/common"
	"sync"
)

var (
	cfg = new(readCfg)
)

const (
	defaultConfigPath = "./"
)

type readCfg struct {
	*config
	locker sync.Mutex
}
type config struct {
	DB *db `yaml:"database",toml:"database"`
}

// TODO
type loadConf interface {
	loadConfig(opts ...opt) *config
}

type defConfig struct {
	filePath string
}
type opt func(*defConfig)

func GetConfig(opts ...opt) *config {
	if cfg.config != nil {
		return cfg.config
	}
	return loadConfig(opts...)
}
func SetConfigFile(filePath string) opt {
	return func(dc *defConfig) {
		dc.filePath = filePath
	}
}

func loadConfig(opts ...opt) *config {
	cfg.locker.Lock()
	defer cfg.locker.Unlock()
	if cfg.config == nil {
		dc := defConfig{
			filePath: defaultConfigPath,
		}

		for _, op := range opts {
			op(&dc)
		}
		err := common.LoadFile(dc.filePath, &cfg.config)
		if err != nil {
			panic(fmt.Sprintf("load system config failed,error msg(%s)", err.Error()))
		}
	}
	return cfg.config
}
