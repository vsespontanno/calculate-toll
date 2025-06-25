package main

import (
	"fmt"

	"github.com/vsespontanno/calculate-toll/types"
)

type Aggregator interface {
	AggregateDistance(Distance types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) *InvoiceAggregator {
	return &InvoiceAggregator{store: store}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("processing and inserting distance in the storage: ", distance)
	err := i.store.Insert(distance)
	if err != nil {
		return err
	}
	return nil
}
