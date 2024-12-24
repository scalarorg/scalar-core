package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
)

// Query labels
const (
	QTokenAddressBySymbol = "token-address-symbol"
	QTokenAddressByAsset  = "token-address-asset"
)

// Bytecode labels
const (
	BCToken  = "token"
	BCBurner = "burner"
)

// Token address labels
const (
	BySymbol  = "symbol"
	ByAsset   = "asset"
	ByAddress = "address"
)

// NewQuerier returns a new querier for the evm module
func NewQuerier(k types.BaseKeeper, n types.Nexus) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		if len(path) <= 1 {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing evm chain name")
		}
		chainKeeper, err := k.ForChain(ctx, exported.ChainName(path[1]))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}

		switch path[0] {
		case QTokenAddressByAsset:
			return QueryTokenAddressByAsset(ctx, chainKeeper, n, path[2])
		case QTokenAddressBySymbol:
			return QueryTokenAddressBySymbol(ctx, chainKeeper, n, path[2])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unknown evm-bridge query endpoint: %s", path[0]))
		}
	}
}

// QueryTokenAddressByAsset returns the address of the token contract by asset
// Deprecated: Use token-info query instead
func QueryTokenAddressByAsset(ctx sdk.Context, k types.ChainKeeper, n types.Nexus, asset string) ([]byte, error) {
	return nil, nil
}

// QueryTokenAddressBySymbol returns the address of the token contract by symbol
// Deprecated: Use token-info query instead
func QueryTokenAddressBySymbol(ctx sdk.Context, k types.ChainKeeper, n types.Nexus, symbol string) ([]byte, error) {
	return nil, nil
}
