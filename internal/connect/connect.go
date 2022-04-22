package connect

import (
	"log"

	"github.com/quankori/go-cosmos/configs"
	"github.com/tendermint/tendermint/rpc/client/http"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func NewRPCClient() *http.HTTP {
	config, _ := configs.LoadConfig()
	client, err := rpchttp.New(config.RpcURI, "/websocket")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Start()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
