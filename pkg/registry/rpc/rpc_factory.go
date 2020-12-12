package rpc

import "context"

type ServiceInfo struct {
	ServiceName string
	Rcvr        interface{}
	Metadata    string
}

type RpcConfig struct {
	// common
	Addr        string
	Network     string
	EtcdAddrs   []string
	BasePath    string
	ServicePath string

	// rpcx
	UpdateInterval int64
}

type RPC interface {
	RegistryService(infos []ServiceInfo)
	CreateClient(servicePath string) RpcClient
	CreateService()
	StartService()
}

type RpcClient interface {
	Call(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error
}

type RpcCreater interface {
	Create(*RpcConfig) RPC
}
