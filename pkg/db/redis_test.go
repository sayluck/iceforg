package db

import (
	"fmt"
	"iceforg/app/test"
	"testing"
	"time"

	rPkg "github.com/go-redis/redis"
)

func init() {
	test.TestInit()
}

func TestGetRedisProvider(t *testing.T) {
	tests := []struct {
		name string
		want *rPkg.Client
	}{
		{
			name: "test_connect",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRedisProvider()
			if got == nil {
				t.Fatal("init redis error")
				return
			}
			pong, err := got.Ping().Result()
			t.Logf("redis ping result:%v,%v", pong, err)
		})
	}
}

func TestSetAndGetVal(t *testing.T) {
	c := GetRedisProvider()
	val := "testVal"
	err := c.Set(val, "100", 60*time.Second).Err()
	if err != nil {
		t.Fatal("set failed," + err.Error())
	}

	v, err := c.Get(val).Result()
	if err != nil {
		t.Fatal("get failed," + err.Error())
	}
	fmt.Println("get val:", v)
}
