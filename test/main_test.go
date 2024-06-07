package test

import (
	"swagger/client"
	"swagger/client/products"
	"testing"
)

func TestNewClient(t *testing.T) {
	c:= client.Default
	params:=products.NewNoParamParams()
	c.Products.NoParam(params)

}
