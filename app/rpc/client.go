package rpc

import (
	"iceforg/pkg/registry/rpc"
	"iceforg/pkg/registry/rpc/rpcx"
)

var RpcClient rpc.RPC

func init() {
	var r rpc.RpcCreater
	// TODO load config
	r = new(rpcx.RpcxCreator)
	RpcClient = r.Create(&rpc.RpcConfig{
		Addr:           "127.0.0.1:60000",
		BasePath:       "iceforg",
		ServicePath:    "user",
		EtcdAddrs:      []string{"192.168.175.1:32790", "192.168.175.1:32788", "192.168.175.1:32786"},
		UpdateInterval: 3600,
	})
}
