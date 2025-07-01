package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vsespontanno/calculate-toll/aggregator/client"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":6060", "listen address")
	aggregatorEndpoint := flag.String("aggregatorEndpoint", "http://localhost:8080", "aggregator endpoint")
	flag.Parse()
	client := client.NewHTTPClient(*aggregatorEndpoint)
	invHandler := NewInvoiceHandler(client)
	http.HandleFunc("/invoice", makeAPIfunc(invHandler.handleGetInvoice))
	logrus.Infof("gateway HTTP motherfucker running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

}

type InvoiceHandler struct {
	client client.Client
}

func NewInvoiceHandler(client client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: client,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	// TODO: fix hardcoded stuff
	inv, err := h.client.GetInvoice(context.Background(), 328764118)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, inv)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func makeAPIfunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(time.Time) {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(time.Now()),
				"uri":  r.RequestURI,
			}).Info("REQ :: ")
		}(time.Now())
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
	}
}
