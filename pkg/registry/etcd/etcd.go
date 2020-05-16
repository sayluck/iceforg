package etcd

import (
	client "github.com/coreos/etcd/clientv3"
)

func GetClient(c *Config) {
	cli, err := client.New(client.Config{
		Endpoints:   c.Addrs,
		DialTimeout: c.DialTimeout, //5 * time.Second
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
}
