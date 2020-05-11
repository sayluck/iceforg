package start

import (
	"iceforg/app/common"
	"iceforg/app/log"
	"iceforg/pkg/config"
	"iceforg/pkg/db"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AppInit() {
	// load config
	cfg := config.GetConfig(
		config.SetConfigFile("F:/goproject/iceforg/resource/config-Files/config.yaml"),
	)

	utils.PrettyJsonPrint(cfg.App)

	// multilingual init
	multilingual.InitMultilingual(
		"F:/goproject/iceforg/resource/config-Files/multilingual_zh.properties")

	// app init
	ginInit(cfg.App)

	// log init
	log.LogInit(cfg.App.Log)

	// database init
	initDB(cfg.DB)
}

func initDB(dbCfg *config.DB) {
	db.InitDB(dbCfg)
}

func ginInit(app *config.App) {
	if strings.ToLower(app.Mode) == common.GinReleaseModle {
		gin.SetMode(gin.ReleaseMode)
	}
}
