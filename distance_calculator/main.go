package main

import (
	"fmt"
	"log"

	"github.com/vsespontanno/calculate-toll/aggregator/client"
)

//type DistanceCalculator struct {
//	consumer DataConsumer
//}

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://localhost:8080"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
	fmt.Println("Working fine motherfucker !")
}
