package main

import (
	"context"
	"fmt"

	"github.com/quankori/go-cosmos/configs"
	"github.com/quankori/go-cosmos/internal/connect"
)

func main() {
	fmt.Println("hello")
	config, _ := configs.LoadConfig()
	var ctx context.Context
	connect.NewRPCClient(config.RpcURI, ctx)
	fmt.Println("exit")
}
