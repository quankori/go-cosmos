package main

import (
	"fmt"

	"github.com/quankori/go-cosmos/pkg/query"
)

func main() {
	fmt.Println("hello")
	query := query.Query{Client: cl, Options: &options}
	balance, err := query.Balances(encodedAddr)
	if err != nil {
		return err
	}
	fmt.Println("exit")
}
