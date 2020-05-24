package test

import (
	"iceforg/pkg/config"
)

func TestInit() {
	cofFilePath := "F:/goproject/iceforg/resource/config-Files/config.yaml"
	config.GetConfig(config.SetConfigFile(cofFilePath))
}
