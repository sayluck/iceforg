package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Products() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.38.135:32777"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "test1"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries
	p.Flush(15 * 1000)
}

//
//import (
//	"github.com/Shopify/sarama"
//	"log"
//	"os"
//	"time"
//)
//
//var (
//	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
//)
//
//func Products() {
//	sarama.Logger = logger
//
//	config := sarama.NewConfig()
//	config.Producer.RequiredAcks = sarama.WaitForAll
//	config.Producer.Partitioner = sarama.NewRandomPartitioner
//	config.Producer.Return.Successes = true
//
//	msg := &sarama.ProducerMessage{}
//	msg.Topic = "test1"
//	msg.Partition = int32(-1)
//	msg.Key = sarama.StringEncoder("key")
//	msg.Value = sarama.ByteEncoder("你好, 世界!")
//	str := []string{"192.168.38.135:32777"}
//	producer, err := sarama.NewSyncProducer(str, config)
//	if err != nil {
//		logger.Println("Failed to produce message: %s", err)
//		os.Exit(500)
//	}
//	defer producer.Close()
//	for {
//		partition, offset, err := producer.SendMessage(msg)
//		if err != nil {
//			logger.Println("Failed to produce message: ", err)
//		}
//		logger.Printf("partition=%d, offset=%d\n", partition, offset)
//		time.Sleep(1 * time.Second)
//	}
//
//}
