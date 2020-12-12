package rpc

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

type Rpc interface {
	RegistryService(infos []ServiceInfo)
	CreateClient(servicePath string) interface{}
	CreateService()
	StartService()
}

type RpcCreater interface {
	Create(*RpcConfig) Rpc
}
