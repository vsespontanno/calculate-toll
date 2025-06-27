package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vsespontanno/calculate-toll/types"
)

type Client struct {
	Endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) AggregateInvoice(distance types.Distance) error {
	b, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Endpoint+"/aggregate", bytes.NewReader(b))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with a non-200 status code: %d", res.StatusCode)
	}
	return nil
}
