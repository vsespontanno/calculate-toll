package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vsespontanno/calculate-toll/types"
)

type HTTPClient struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

func (c *HTTPClient) GetInvoice(ctx context.Context, id int) (*types.Invoice, error) {
	invReq := types.GetInvoiceRequest{
		ObuID: int32(id),
	}
	b, err := json.Marshal(&invReq)
	if err != nil {
		return nil, err
	}
	logrus.Printf("req to %s\n", c.Endpoint+"/invoice"+fmt.Sprintf("?id=%d", id))
	req, err := http.NewRequest("POST", c.Endpoint+"/invoice"+fmt.Sprintf("?obu=%d", id), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the service responded with a non-200 status code: %d", res.StatusCode)
	}
	var inv types.Invoice
	err = json.NewDecoder(res.Body).Decode(&inv)
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

func (c *HTTPClient) Aggregate(ctx context.Context, aggReq *types.AggregateRequest) error {
	b, err := json.Marshal(aggReq)
	if err != nil {
		return err
	}
	logrus.Printf("posting req to %s\n", c.Endpoint+"/aggregate")
	req, err := http.NewRequest("POST", c.Endpoint+"/aggregate", bytes.NewReader(b))
	if err != nil {
		logrus.Errorf("having fucking error which is stage 1: %v", err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("having fucking error which is stage 2: %v", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with a non-200 status code: %d", res.StatusCode)
	}
	return nil
}
