package etcd

import "time"

type Config struct {
	Addrs       []string
	DialTimeout time.Duration
}
