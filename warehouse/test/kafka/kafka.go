package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.38.135:32777",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"test1", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}

//
//import (
//	"fmt"
//	"github.com/Shopify/sarama"
//	"strings"
//	"sync"
//)
//
//var (
//	wg sync.WaitGroup
//)
//
//func Consumer() {
//
//	sarama.Logger = logger
//
//	consumer, err := sarama.NewConsumer(strings.Split("192.168.38.135:32777", ","), nil)
//	if err != nil {
//		logger.Println("Failed to start consumer: %s", err)
//	}
//	for {
//		partitionList, err := consumer.Partitions("test1")
//		if err != nil {
//			logger.Println("Failed to get the list of partitions: ", err)
//		}
//
//		for partition := range partitionList {
//			pc, err := consumer.ConsumePartition("test1", int32(partition), sarama.OffsetNewest)
//			if err != nil {
//				logger.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
//			}
//			defer pc.AsyncClose()
//
//			wg.Add(1)
//
//			go func(sarama.PartitionConsumer) {
//				defer wg.Done()
//				for msg := range pc.Messages() {
//					fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
//					fmt.Println()
//				}
//			}(pc)
//		}
//
//		wg.Wait()
//
//		logger.Println("Done consuming topic test1")
//	}
//
//	consumer.Close()
//}
