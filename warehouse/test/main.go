package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	md5Hash := md5.New()
	fmt.Println(hex.EncodeToString(md5Hash.Sum([]byte("hello"))))
	fmt.Println(string(md5Hash.Sum([]byte("hello"))))
	time.Sleep(time.Second * 1)
}
