package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewAddressValidator returns the callback for validating BTC addresses
func NewAddressValidator() nexus.AddressValidator {
	return func(ctx sdk.Context, address nexus.CrossChainAddress) error {
		// TODO: validate btc address
		clog.Red("TODO: validate address", address)

		// chain := address.Chain.Name // evm|11155111, bitcoin|1, bitcoin|4

		// if !common.IsHexAddress(address.Address) {
		// 	return fmt.Errorf("not an hex address")
		// }
		return nil
	}
}
