package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &ReserveRedeemUtxoRequest{}

func NewReserveRedeemUtxoRequest(sender sdk.AccAddress, chain string, address string, symbol string, amount uint64) *ReserveRedeemUtxoRequest {
	return &ReserveRedeemUtxoRequest{
		Sender:  sender,
		Chain:   nexus.ChainName(chain),
		Address: address,
		Symbol:  symbol,
		Amount:  amount,
	}
}

func (msg *ReserveRedeemUtxoRequest) ValidateBasic() error {
	// TODO: validate the txID
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}

	if nexus.ChainFamily(msg.Chain) != nexus.EVM {
		return fmt.Errorf("chain %s is not a EVM chain", msg.Chain)
	}

	if msg.Address == "" {
		return fmt.Errorf("address is required")
	}

	if msg.Symbol == "" {
		return fmt.Errorf("symbol is required")
	}

	if msg.Amount == 0 {
		return fmt.Errorf("amount is required")
	}

	return nil
}

func (msg *ReserveRedeemUtxoRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
