package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (client *BtcClient) ProcessDestinationTxsConfirmation(event *types.EventConfirmDestTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })

	clog.Blue("[BTC] ProcessDestinationTxsConfirmation", "txIDs", txIDs)

	return nil, fmt.Errorf("not implemented")
}
