package main

import (
	"fmt"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	libclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)

func main() {
	fmt.Println("ok")
	NewRPCClient("https://rpc.titus-1.archway.tech")
	fmt.Println("ok 2")
}

// func queryState() error {
// 	// myAddress, err := sdk.AccAddressFromBech32("archway1fs8lczpfms4gp7vy0kcmyrx607gxgz8suel4w6")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// Create a connection to the gRPC server.
// 	grpcConn, _ := grpc.Dial(
// 		"https://rpc.titus-1.archway.tech", // your gRPC server address.
// 		grpc.WithInsecure(),                // The SDK doesn't support any transport security mechanism.
// 	)
// 	defer grpcConn.Close()
// 	// This creates a gRPC client to query the x/bank service.
// 	bankClient := banktypes.NewQueryClient(grpcConn)
// 	fmt.Println(bankClient)
// 	bankRes, err := bankClient.Balance(
// 		context.Background(),
// 		&banktypes.QueryBalanceRequest{Address: "archway1fs8lczpfms4gp7vy0kcmyrx607gxgz8suel4w6"},
// 	)
// 	fmt.Println(err)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(bankRes.GetBalance()) // Prints the account balance

// 	return nil
// }

func NewRPCClient(addr string) (*rpchttp.HTTP, error) {
	httpClient, err := libclient.DefaultHTTPClient(addr)
	if err != nil {
		return nil, err
	}
	fmt.Println(httpClient.Get("health"))
	// httpClient.Timeout = timeout
	rpcClient, err := rpchttp.NewWithClient(addr, "/websocket", httpClient)
	if err != nil {
		return nil, err
	}
	fmt.Println(rpcClient)
	return rpcClient, nil
}
