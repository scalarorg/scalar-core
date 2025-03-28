package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &ReserveRedeemUtxoRequest{}

func NewReserveRedeemUtxoRequest(sender sdk.AccAddress, sourceChain, destChain nexus.ChainName, address string, symbol string, amount uint64) *ReserveRedeemUtxoRequest {
	return &ReserveRedeemUtxoRequest{
		Sender:      sender,
		SourceChain: sourceChain,
		DestChain:   destChain,
		Address:     address,
		Symbol:      symbol,
		Amount:      amount,
	}
}

func (msg *ReserveRedeemUtxoRequest) ValidateBasic() error {
	// TODO: validate the txID
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.SourceChain == "" {
		return fmt.Errorf("source chain is required")
	}
	sourceChainName := nexus.ChainName(msg.SourceChain)
	if sourceChainName.GetFamily() != nexus.EVM {
		return fmt.Errorf("source chain %s is not a EVM chain", msg.SourceChain)
	}

	if msg.DestChain == "" {
		return fmt.Errorf("dest chain is required")
	}
	destChainName := nexus.ChainName(msg.DestChain)
	if destChainName.GetFamily() != nexus.BITCOIN {
		return fmt.Errorf("dest chain %s is not a BTC chain", msg.DestChain)
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
