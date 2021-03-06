package connect

import (
	"log"

	"github.com/quankori/go-cosmos/configs"
	"github.com/tendermint/tendermint/rpc/client/http"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func NewRPCSocket() *http.HTTP {
	config, _ := configs.LoadConfig()
	client, err := rpchttp.New(config.RpcURI + "/websocket")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewRPCClient() *http.HTTP {
	config, _ := configs.LoadConfig()
	client, err := rpchttp.New(config.TendermintURI)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
