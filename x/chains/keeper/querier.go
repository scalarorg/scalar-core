package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/evm/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier returns a new querier for the evm module
func NewQuerier(k types.BaseKeeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		// TODO: implement
		return nil, nil
	}
}
