package main

import (
	"context"
	"fmt"
	"iceforg/modules/userCenter"
	"iceforg/pkg/registry/rpc"
	"iceforg/pkg/registry/rpc/rpcx"
	"iceforg/pkg/utils"
)

func main() {
	var r rpc.RpcCreater
	r = new(rpcx.RpcxCreator)
	c := r.Create(&rpc.RpcConfig{
		Addr:           "127.0.0.1:60000",
		BasePath:       "iceforg",
		ServicePath:    "user",
		EtcdAddrs:      []string{"192.168.175.1:32790", "192.168.175.1:32788", "192.168.175.1:32786"},
		UpdateInterval: 3600,
	})

	xClient := c.CreateClient("user")

	reply := &userCenter.UserDetail{}
	err := xClient.Call(context.Background(), "Detail", "hws", reply)

	if err != nil {
		utils.PrettyJsonPrint(err.Error())
		return
	}
	fmt.Println("get data ========== ")
	utils.PrettyJsonPrint(reply)
	fmt.Println("get data ========== ")
}
