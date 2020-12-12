package main

import (
	"iceforg/modules/common"
	"iceforg/modules/userCenter"
	"iceforg/pkg/config"
	"iceforg/pkg/db"
	"iceforg/pkg/registry/rpc"
	"iceforg/pkg/registry/rpc/rpcx"
	"iceforg/pkg/utils"
)

func main() {
	initService()

	// user rpcx
	r := new(rpcx.RpcxCreator)
	typeRpcx := r.Create(getRpcConf())

	common.IceLog.Debug("Start Create Rpc Service....")
	typeRpcx.CreateService()
	typeRpcx.RegistryService([]rpc.ServiceInfo{
		registryUserCenter("usercenter"),
	})
	common.IceLog.Debug("Start Rpc Service....")
	typeRpcx.StartService()
}

func getRpcConf() *rpc.RpcConfig {
	return &rpc.RpcConfig{
		Addr:           "127.0.0.1:60000",
		Network:        "tcp",
		BasePath:       "iceforg",
		EtcdAddrs:      []string{"192.168.175.1:32790", "192.168.175.1:32788", "192.168.175.1:32786"},
		UpdateInterval: 3600,
	}
}

func registryUserCenter(serviceName string) rpc.ServiceInfo {
	info := rpc.ServiceInfo{
		ServiceName: serviceName,
		Rcvr:        new(userCenter.UserCenter),
		Metadata:    "",
	}
	return info
}

func initService() {
	// load config
	cfg := config.GetConfig(
		config.SetConfigFile("F:/goproject/iceforg/resource/config-files/config.yaml"),
	)

	// log init
	common.LogInit()
	common.SetLogConfig(cfg.App.Log)
	common.IceLog.Debugf("load app config:%s", utils.PrettyJsonPrint(cfg.App))

	// database init
	initDB(cfg.DB)
}

func initDB(dbCfg *config.DB) {
	db.InitDB(dbCfg)
}
