package main

import (
	"fmt"

	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
	"github.com/what-the-func/hello-loom/types"
)

// HelloWorld contract
type HelloWorld struct{}

// Meta data
func (c *HelloWorld) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "helloworld",
		Version: "1.0.0",
	}, nil
}

// Init sets up the contract like a contructor
func (c *HelloWorld) Init(ctx contract.Context, req *types.HelloRequest) error {
	return nil
}

// Hello method
func (c *HelloWorld) Hello(ctx contract.StaticContext, req *types.HelloRequest) (*types.HelloResponse, error) {
	return &types.HelloResponse{
		Out: fmt.Sprintf("Hello, %s", req.In),
	}, nil
}

// Contract instance
var Contract plugin.Contract = contract.MakePluginContract(&HelloWorld{})

func main() {
	plugin.Serve(Contract)
}
