package query

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/quankori/go-cosmos/internal/connect"
)

func SubscribeEvent() {
	client := connect.NewRPCClient()
	cw20Contract := "juno1suhgf5svhu4usrurvxzlgn54ksxmn8gljarjtxqnapv8kjnp4nrsf8smqw"
	err := client.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	query := fmt.Sprintf("message.module='wasm' AND wasm._contract_address='%s' AND wasm.action='transfer'", cw20Contract)
	txs, err := client.Subscribe(ctx, "wasm-client", query)
	if err != nil {
		log.Fatal(err)
	}

	for e := range txs {
		from := e.Events["wasm.from"][0]
		to := e.Events["wasm.to"][0]
		amount := e.Events["wasm.amount"][0]

		fmt.Println("Tx: " + e.Events["tx.hash"][0])
		fmt.Printf("Sender: %s \nRecipient: %s \nAmount: %s \n", from, to, amount)
	}
}
