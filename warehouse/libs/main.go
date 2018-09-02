package main

import (
	log "github.com/ctlog"
	"github.com/iceforg/warehouse/libs/docker"
)

func ConfigInit() {

	// log init
	log.SetLogLevel("info")
	log.SetLogDir("/tmp/log", "iceforg")

	// Docker client init
	docker.NewClient()
}

func main() {
	ConfigInit()

	log.Infoln("================= CONFIG INIT OK ================")
	run()
	log.Infoln("==================== STARTING ===================")
}

func run() {

}
