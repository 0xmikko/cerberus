package main

import (
	"fmt"
	"github.com/linkpoolio/bridges"
	"net/http"
)

type Cerberus struct{}

// Run is the bridge.Bridge Run implementation that returns the confirmation response
func (cc *Cerberus) Run(h *bridges.Helper) (interface{}, error) {
	r := make(map[string]interface{})

	id := h.Data.Get("id").String()
	requestURL := fmt.Sprintf("https:/cerberus.ledger-labs.com/adapter/transactions/%s/", id)

	err := h.HTTPCall(
		http.MethodGet,
		requestURL,
		&r,
	)
	return r["confirmation"], err
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
