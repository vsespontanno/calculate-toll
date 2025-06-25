package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"github.com/vsespontanno/calculate-toll/types"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string, svc CalculatorServicer) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	return &KafkaConsumer{consumer: c, calcService: svc}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Starting Kafka consumer")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consume reading message error : %v", err)
			continue
		}
		var data types.OBUData
		if err = json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON ser error : %v", err)
			continue
		}
		dist, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("Calculate distance error : %v", err)
			continue
		}
		fmt.Println(dist)

	}
}
