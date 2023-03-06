package wasmbinding

import (
	denomkeeper "github.com/Team-Kujira/core/x/denom/keeper"

	oraclekeeper "github.com/Team-Kujira/core/x/oracle/keeper"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
)

type QueryPlugin struct {
	denomKeeper  denomkeeper.Keeper
	bankkeeper   bankkeeper.Keeper
	oraclekeeper oraclekeeper.Keeper
	ibckeeper    ibckeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(bk bankkeeper.Keeper, ok oraclekeeper.Keeper, dk denomkeeper.Keeper, ik ibckeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		denomKeeper:  dk,
		bankkeeper:   bk,
		oraclekeeper: ok,
		ibckeeper:    ik,
	}
}
