package start

import (
	"iceforg/app/common"
	"iceforg/app/log"
	"iceforg/app/validate"
	"iceforg/pkg/config"
	"iceforg/pkg/db"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AppInit() {
	// log init
	log.LogInit()

	defer initRecover()

	// load config
	cfg := config.GetConfig(
		config.SetConfigFile("F:/goproject/iceforg/resource/config_files/config.yaml"),
	)

	// log config
	log.SetLogConfig(cfg.App.Log)
	log.Log.Debugf("load app config:%s", utils.PrettyJsonPrint(cfg.App))

	// multilingual init
	multilingual.InitMultilingual(
		"F:/goproject/iceforg/resource/config_files/multilingual_zh.properties")

	// app init
	ginInit(cfg.App)

	// database init
	initDB(cfg.DB)

	validate.InitValidate()
}

func initDB(dbCfg *config.DB) {
	db.InitDB(dbCfg)
}

func ginInit(app *config.App) {
	if strings.ToLower(app.Mode) == common.GinReleaseModle {
		gin.SetMode(gin.ReleaseMode)
	}
}

func initRecover() {
	if err := recover(); err != nil {
		log.Log.Fatalf("app init failed:%v\n", err)
	}
}
