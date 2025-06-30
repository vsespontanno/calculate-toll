package client

import (
	"context"
	"fmt"

	"github.com/vsespontanno/calculate-toll/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	client   types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error while making new client: ", err)
		return nil, err
	}
	c := types.NewAggregatorClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client:   c,
	}, nil
}

func (c *GRPCClient) Aggregate(ctx context.Context, req *types.AggregateRequest) error {
	_, err := c.client.Aggregate(ctx, req)
	return err
}
