package test

import (
	"iceforg/pkg/config"
	"iceforg/pkg/db"
)

func TestInit() {
	cofFilePath := "F:/goproject/iceforg/resource/config-Files/config.yaml"
	cof := config.GetConfig(config.SetConfigFile(cofFilePath))

	db.InitDB(cof.DB)
}
