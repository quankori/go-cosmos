package main

import (
	"context"
	"fmt"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	libclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)

func main() {
	fmt.Println("ok")
	var ctx context.Context
	NewRPCClient("https://rpc.cosmos.network", ctx)
	// NewRPCClient("https://rpc.titus-1.archway.tech")
	fmt.Println("ok 2")
}

func NewRPCClient(addr string, ctx context.Context) (*rpchttp.HTTP, error) {
	httpClient, err := libclient.DefaultHTTPClient(addr)
	if err != nil {
		return nil, err
	}

	response, err := rpchttp.NewWithClient(addr, "/status", httpClient)
	fmt.Println(response.Status(ctx))
	// httpClient.Timeout = timeout
	rpcClient, err := rpchttp.NewWithClient(addr, "/websocket", httpClient)
	if err != nil {
		return nil, err
	}
	fmt.Println(rpcClient)
	return rpcClient, nil
}
