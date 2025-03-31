package xchain

import (
	"context"
	"fmt"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	"github.com/scalarorg/scalar-core/utils/log"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	"github.com/scalarorg/scalar-core/x/chains/types"
	cov "github.com/scalarorg/scalar-core/x/covenant/types"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

// Manager manages all communication with Ethereum
type Manager struct {
	rpcs                      map[chain.ChainInfoBytes]xcommon.Client
	broadcaster               broadcast.Broadcaster
	validator                 sdk.ValAddress
	proxy                     sdk.AccAddress
	latestFinalizedBlockCache xcommon.LatestFinalizedBlockCache
}

// NewManager returns a new Manager instance
func NewManager(
	clientCtx sdkClient.Context,
	rpcs map[chain.ChainInfoBytes]xcommon.Client,
	broadcaster broadcast.Broadcaster,
	valAddr sdk.ValAddress,
) *Manager {
	return &Manager{
		rpcs:        rpcs,
		broadcaster: broadcaster,
		validator:   valAddr,
		proxy:       clientCtx.FromAddress,
	}
}

func (mgr Manager) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted) error {
	if !mgr.isParticipantOf(event.Participants) {
		pollIDs := slices.Map(event.PollMappings, func(m types.PollMapping) vote.PollID { return m.PollID })
		mgr.logger("poll_ids", pollIDs).Debug("ignoring staking txs confirmation poll: not a participant")
		return nil
	}

	mgr.logger("event", event).Debug("processing staking txs confirmation poll")

	chainInfoBytes := chain.ChainInfoBytes{}

	err := chainInfoBytes.FromString(event.Chain.String())
	if err != nil {
		return err
	}

	client, ok := mgr.rpcs[chainInfoBytes]
	if !ok {
		return fmt.Errorf("rpc client not found for chain %s", event.Chain.String())
	}

	votes, err := client.ProcessSourceTxsConfirmation(event, mgr.proxy)
	if err != nil {
		return err
	}

	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)

	return err
}

func (mgr Manager) ProcessRedeemTxConfirmation(event *cov.ConfirmRedeemTxStarted) error {
	if !mgr.isParticipantOf(event.Participants) {
		mgr.logger("poll_id", event.PollID).Debug("ignoring redeem tx confirmation poll: not a participant")
		return nil
	}

	mgr.logger("event", event).Debug("processing redeem tx confirmation poll")

	chainInfoBytes := chain.ChainInfoBytes{}

	err := chainInfoBytes.FromString(event.Chain.String())
	if err != nil {
		return err
	}

	btcClient, ok := mgr.rpcs[chainInfoBytes].(xcommon.BtcClient)
	if !ok {
		return fmt.Errorf("rpc client not found for chain %s", event.Chain.String())
	}

	votes, err := btcClient.ProcessRedeemTxsConfirmation(event, mgr.proxy)
	if err != nil {
		return err
	}
	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)
	return err
}

func (mgr Manager) ProcessSwitchedPhaseConfirmation(event *cov.ConfirmSwitchedPhaseStarted) error {
	if !mgr.isParticipantOf(event.Participants) {
		mgr.logger("poll_id", event.PollID).Debug("ignoring switched phase confirmation poll: not a participant")
		return nil
	}

	mgr.logger("event", event).Debug("processing switched phase confirmation poll")

	chainInfoBytes := chain.ChainInfoBytes{}

	err := chainInfoBytes.FromString(event.Chain.String())
	if err != nil {
		return err
	}

	evmClient, ok := mgr.rpcs[chainInfoBytes].(xcommon.EvmClient)
	if !ok {
		return fmt.Errorf("rpc client not found for chain %s", event.Chain.String())
	}

	votes, err := evmClient.ProcessSwitchedPhaseConfirmation(event, mgr.proxy)
	if err != nil {
		return err
	}
	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)
	if err != nil {
		log.Errorf("[Manager] poll %s failed to broadcast vote: %++v", event.PollID.String(), err)
		return err
	}
	return nil
}

// isParticipantOf checks if the validator is in the poll participants list
func (mgr Manager) isParticipantOf(participants []sdk.ValAddress) bool {
	return slices.Any(participants, func(v sdk.ValAddress) bool { return v.Equals(mgr.validator) })
}

func (mgr Manager) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"listener", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}
