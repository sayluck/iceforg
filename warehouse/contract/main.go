package main

import (
	"github.com/ctlog"
	"github.com/iceforg/warehouse/contract/docker"
)

func ConfigInit() {
	ctlog.SetLogLevel("info")
	ctlog.SetLogDir("/tmp/log", "iceforg")
}

func main() {
	ConfigInit()

	// get docker client
	docker.GetDockerClient()
	ctlog.Infoln("================= STARTING ================")
	ctlog.Debugf("================= STARTING ================%s", "asd")

	//

}
