package app

import (
	"iceforg/app/controler"
	"iceforg/pkg/config"
)

func Run() {
	r := controler.Route{App: config.GetConfig().App}
	r.Router()
}
