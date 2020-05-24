package db

import (
	"crypto/tls"
	"fmt"
	"iceforg/pkg/config"
	"sync"
	"time"

	etcdCli "github.com/coreos/etcd/clientv3"
)

var etcdProvider = new(etcd)

type etcd struct {
	config *config.Etcd
	client *etcdCli.Client
	locker sync.Mutex
}

func GetEtcdProvider() *etcdCli.Client {
	if etcdProvider.client == nil {
		etcdProvider.initer()
	}
	return etcdProvider.client
}

func (e *etcd) initer() {
	if e.client == nil {
		e.locker.Lock()
		defer e.locker.Unlock()
		var err error
		e.client, err = etcdCli.New(e.etcdOpt())
		if err != nil {
			panic(fmt.Sprintf("get etcd client failed,%v", err))
		}
	}
}

func (e *etcd) etcdOpt() etcdCli.Config {

	cnf := etcdCli.Config{
		Endpoints:   e.config.Addrs,
		DialTimeout: e.config.DialTimeout * time.Second,
	}
	if e.config.Secure || e.config.TLSConfig != nil {
		tlsConfig := e.config.TLSConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		cnf.TLS = tlsConfig
	}
	return cnf
}
