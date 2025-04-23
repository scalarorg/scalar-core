package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &ConfirmRedeemTxsRequest{}

func NewConfirmRedeemTxRequest(sender sdk.AccAddress, chain string, groupUid string, txIDs []chains.Hash) *ConfirmRedeemTxsRequest {
	custodianGroupUid, err := chains.HashFromHex(groupUid)
	if err != nil {
		log.Error().Err(err).Msg("Invalid groupUid")
	}
	return &ConfirmRedeemTxsRequest{
		Sender:            sender,
		Chain:             nexus.ChainName(chain),
		CustodianGroupUID: custodianGroupUid,
		TxIDs:             txIDs,
	}
}

func (msg *ConfirmRedeemTxsRequest) ValidateBasic() error {
	// TODO: validate the txID
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}
	if len(msg.TxIDs) == 0 {
		return fmt.Errorf("txIDs are required")
	}
	return nil
}

func (msg *ConfirmRedeemTxsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
