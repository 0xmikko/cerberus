package main

import (
	"fmt"
	"github.com/linkpoolio/bridges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCerberus_Run(t *testing.T) {
	wa := Cerberus{}

	data := make(map[string]interface{})
	data["id"] = "0526d68c-135d-46ca-a70a-8f153e9e0416"

	json, err := bridges.ParseInterface(data)
	h := bridges.NewHelper(json)
	val, err := wa.Run(h)

	fmt.Println(err)
	confirmation, ok := val.(bool)
	assert.True(t, ok)
	assert.True(t, confirmation == true)
	assert.Nil(t, err)
}

func TestCerberus_Opts(t *testing.T) {
	cc := Cerberus{}
	opts := cc.Opts()
	assert.Equal(t, opts.Name, "Cerberus")
	assert.True(t, opts.Lambda)
}
