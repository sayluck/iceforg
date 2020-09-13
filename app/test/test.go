package test

import (
	"iceforg/pkg/config"
)

func TestInit() {
	cofFilePath := "F:/goproject/iceforg/resource/config_files/config.yaml"
	config.GetConfig(config.SetConfigFile(cofFilePath))
}
