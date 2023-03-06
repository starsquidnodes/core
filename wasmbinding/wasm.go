package wasmbinding

import (
	denomkeeper "github.com/Team-Kujira/core/x/denom/keeper"

	oraclekeeper "github.com/Team-Kujira/core/x/oracle/keeper"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	bankkeeper "github.com/terra-money/alliance/custom/bank/keeper"
)

func RegisterCustomPlugins(
	bank bankkeeper.Keeper,
	oracle oraclekeeper.Keeper,
	denom denomkeeper.Keeper,
	auth authkeeper.AccountKeeper,
	ibc ibckeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(bank, oracle, denom, ibc)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})

	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, denom, auth),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
