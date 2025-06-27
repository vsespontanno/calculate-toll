package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vsespontanno/calculate-toll/types"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{next: next}
}
func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("aggregated distance")

	}(time.Now())
	return m.next.AggregateDistance(distance)
}

func (m *LogMiddleware) CalculateInvoice(obuID int) (inv *types.Invoice, err error) {

	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":     time.Since(start),
			"err":      err,
			"obuID":    obuID,
			"distance": distance,
			"amount":   amount,
		}).Info("calculateInvoice")
	}(time.Now())
	inv, err = m.next.CalculateInvoice(obuID)
	return
}
