package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func main() {
	fmt.Println("ok")
	queryState()
	fmt.Println("ok 2")
}

func queryState() error {
	// myAddress, err := sdk.AccAddressFromBech32("archway1fs8lczpfms4gp7vy0kcmyrx607gxgz8suel4w6")
	// if err != nil {
	// 	return err
	// }

	// Create a connection to the gRPC server.
	grpcConn, _ := grpc.Dial(
		"https://rpc.cosmos.network", // your gRPC server address.
		grpc.WithInsecure(),          // The SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: "cosmos1xyjlrf7fq9fg4jjqgk8f2vx3a4mavfx3g8a5x8"},
	)
	fmt.Println(bankRes)
	if err != nil {
		return err
	}

	fmt.Println(bankRes.GetBalance()) // Prints the account balance

	return nil
}
