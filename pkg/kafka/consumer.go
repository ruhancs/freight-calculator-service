package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string, servers string, msgChan chan *ckafka.Message) {
	kafkaConsumer,err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id": "freigth_service",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	kafkaConsumer.SubscribeTopics(topics,nil)

	for {
		msg,err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}