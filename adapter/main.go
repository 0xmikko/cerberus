package main

import (
	"github.com/linkpoolio/bridges"
	"net/http"
)

// Cerberus is the most basic Bridge implementation, as it only calls the api:
// https://min-ap.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,JPY,EUR
type Cerberus struct{}

// Run is the bridge.Bridge Run implementation that returns the price response
func (cc *Cerberus) Run(h *bridges.Helper) (interface{}, error) {
	r := make(map[string]interface{})
	err := h.HTTPCall(
		http.MethodGet,
		"https:/cerberus.ledger-labs.com/api/transaction/",
		&r,
	)
	return r, err
}

// Opts is the bridge.Bridge implementation
func (cc *Cerberus) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "Cerberus",
		Lambda: true,
	}
}

func main() {
	bridges.NewServer(&Cerberus{}).Start(8080)
}
