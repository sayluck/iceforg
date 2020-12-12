package db

import (
	"context"
	"fmt"
	"iceforg/app/test"
	"iceforg/pkg/config"
	"iceforg/pkg/utils"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func TestGetClient(t *testing.T) {
	test.TestInit()
	InitDB(config.GetConfig().DB)
	cli := GetEtcdProvider()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	resp1, err := cli.Put(ctx, "key1", "for test")
	if err != nil {
		t.Fatalf("put etcd client error,%v", err)
	}
	utils.PrettyJsonPrint(resp1)

	resp, err := cli.Get(ctx, "key1",
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		t.Fatalf("get etcd client error,%v", err)
	}
	fmt.Printf("%v", resp.Kvs)
	cancel()
}
