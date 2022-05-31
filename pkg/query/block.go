package query

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/quankori/go-cosmos/internal/connect"
	"github.com/tendermint/tendermint/rpc/coretypes"
	"github.com/tendermint/tendermint/types"
)

func SubscribeBlock() {
	client := connect.NewRPCSocket()
	err := client.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// query := "tm.event = 'Tx'"
	query := "tm.event = 'NewBlock'"
	txs, err := client.Subscribe(ctx, "test-client", query)
	if err != nil {
		log.Fatal(err)
	}
	for e := range txs {
		switch data := e.Data.(type) {
		case types.EventDataNewBlock:
			fmt.Printf("Block %s - Height: %d \n", hex.EncodeToString(data.Block.Hash()), data.Block.Height)
			break
		}
	}
}

func BlockResults() (*coretypes.ResultBlockResults, error) {
	client := connect.NewRPCClient()
	var n int64 = 905719
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := client.BlockResults(ctx, &n)
	return res, err
}
