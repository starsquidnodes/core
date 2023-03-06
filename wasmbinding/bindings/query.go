package bindings

import (
	denom "github.com/Team-Kujira/core/x/denom/wasm"
	oracle "github.com/Team-Kujira/core/x/oracle/wasm"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

// DenomQuery contains denom custom queries.

type CosmosQuery struct {
	Denom  *denom.DenomQuery
	Bank   *BankQuery
	Oracle *oracle.OracleQuery
	IBC    *IBCQuery
}

type BankQuery struct {
	DenomMetadata *banktypes.QueryDenomMetadataRequest `json:"denom_metadata,omitempty"`
	Supply        *banktypes.QuerySupplyOfRequest      `json:"supply,omitempty"`
}

type IBCQuery struct {
	Verify *IBCQueryVerify `json:"verify,omitempty"`
}

type IBCQueryVerify struct {
	ConnectionId string           `json:"connection_id,omitempty"`
	ChainId      string           `json:"chain_id,omitempty"`
	Height       int64            `json:"height,omitempty"`
	Module       string           `json:"module,omitempty"`
	Request      []byte           `json:"request,omitempty"`
	Result       []byte           `json:"result,omitempty"`
	ProofOps     *crypto.ProofOps `json:"proof_ops,omitempty"`
}
