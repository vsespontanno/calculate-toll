package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vsespontanno/calculate-toll/types"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{next: next}
}
func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("Producing data to Kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
