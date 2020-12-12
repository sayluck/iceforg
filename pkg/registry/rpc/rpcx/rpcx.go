package rpcx

import (
	"iceforg/pkg/registry/rpc"
	"time"

	"github.com/smallnest/rpcx/client"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

type RpcxCreator struct {
}

func (r *RpcxCreator) Create(c *rpc.RpcConfig) rpc.Rpc {
	return &Rpcx{
		Addr:           c.Addr,
		Network:        c.Network,
		EtcdAddrs:      c.EtcdAddrs,
		BasePath:       c.BasePath,
		ServicePath:    c.ServicePath,
		UpdateInterval: c.UpdateInterval,
	}
}

type Rpcx struct {
	Addr           string
	Network        string
	EtcdAddrs      []string
	BasePath       string
	ServicePath    string
	UpdateInterval int64
	*server.Server
	client.XClient
}

func (r *Rpcx) RegistryService(infos []rpc.ServiceInfo) {
	for _, v := range infos {
		err := r.RegisterName(v.ServiceName, v.Rcvr, v.Metadata)
		if err != nil {
			panic("registry service failed:" + err.Error())
		}
	}
}
func (r *Rpcx) CreateClient(serviceName string) interface{} {
	d := client.NewEtcdV3Discovery(r.BasePath, r.ServicePath, r.EtcdAddrs, nil)
	r.XClient = client.NewXClient(serviceName,
		client.Failtry, client.RandomSelect, d, client.DefaultOption)
	return r.XClient
}
func (r *Rpcx) CreateService() {
	r.Server = server.NewServer()
	etcdReg := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: r.Network + "@" + r.Addr,
		EtcdServers:    r.EtcdAddrs,
		BasePath:       r.BasePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Duration(r.UpdateInterval),
	}
	if err := etcdReg.Start(); err != nil {
		panic("create rpcx service failed:" + err.Error())
	}
	r.Server.Plugins.Add(etcdReg)
}

func (r *Rpcx) StartService() {
	if err := r.Serve(r.Network, r.Addr); err != nil {
		panic("start rpcx service faild:" + err.Error())
	}
}
