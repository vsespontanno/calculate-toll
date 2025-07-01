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

	httpClient := client.NewHTTPClient(aggregatorEndpoint)
	// grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
	fmt.Println("Working fine motherfucker !")
}
