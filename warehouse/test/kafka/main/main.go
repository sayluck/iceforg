package main

import (
	"github.com/iceforg/warehouse/test/kafka"
	"time"
)

func main() {
	go kafka.Products()

	go kafka.Consumer()
	time.Sleep(100 * time.Minute)
}
