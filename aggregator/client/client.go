package client

import (
	"context"

	"github.com/vsespontanno/calculate-toll/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
