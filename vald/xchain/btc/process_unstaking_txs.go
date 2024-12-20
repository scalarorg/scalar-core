package btc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

func (client *BtcClient) ProcessUnstakingTxsConfirmation(event *types.EventConfirmUnstakingTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })

	clog.Blue("ProcessUnstakingTxsConfirmation", "txIDs", txIDs)

	return nil, nil
}
