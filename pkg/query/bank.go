package query

import (
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BalanceWithAddressRPC returns the balance of all coins for a single account.
func BalanceWithAddressRPC(q *Query, address string) (*bankTypes.QueryAllBalancesResponse, error) {
	req := &bankTypes.QueryAllBalancesRequest{Address: address, Pagination: q.Options.Pagination}
	queryClient := bankTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.AllBalances(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// TotalSupplyRPC returns the supply of all coins
func TotalSupplyRPC(q *Query) (*bankTypes.QueryTotalSupplyResponse, error) {
	req := &bankTypes.QueryTotalSupplyRequest{Pagination: q.Options.Pagination}
	queryClient := bankTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.TotalSupply(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
