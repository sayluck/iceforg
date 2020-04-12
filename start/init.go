package start

import (
	"fmt"
	"iceforg/pkg/config"
	"iceforg/pkg/db"
)

func AppInit() {
	cfg := config.GetConfig(
		config.SetConfigFile("F:/goproject/iceforg/resource/config-Files/config.yaml"),
	)
	fmt.Printf("config(%+v)\n", cfg)
	initDB(cfg.DB)
}

func initDB(dbCfg *config.DB) {
	db.InitDB(dbCfg)
}
