package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/x/nexus/exported"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
)

// NewGenesisState is the constructor of GenesisState
func NewGenesisState(
	params Params,
	nonce uint64,
	chains []exported.Chain,
	chainStates []ChainState,
	linkedAddresses []LinkedAddresses,
	transfers []exported.CrossChainTransfer,
	fee exported.TransferFee,
	feeInfos []exported.FeeInfo,
	rateLimits []RateLimit,
	transferEpochs []TransferEpoch,
	messages []exported.GeneralMessage,
	messageNonce uint64,
) *GenesisState {
	return &GenesisState{
		Params:          params,
		Nonce:           nonce,
		Chains:          chains,
		ChainStates:     chainStates,
		LinkedAddresses: linkedAddresses,
		Transfers:       transfers,
		Fee:             fee,
		FeeInfos:        feeInfos,
		RateLimits:      rateLimits,
		TransferEpochs:  transferEpochs,
		Messages:        messages,
		MessageNonce:    messageNonce,
	}
}

// DefaultGenesisState creates the default genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
		0,
		[]exported.Chain{chains.Ethereum, scalarnet.Scalarnet, chains.Bitcoin},
		[]ChainState{{
			Chain:  scalarnet.Scalarnet,
			Assets: []exported.Asset{exported.NewAsset(scalarnet.NativeAsset, true)},
		}},
		[]LinkedAddresses{},
		[]exported.CrossChainTransfer{},
		exported.TransferFee{},
		[]exported.FeeInfo{},
		[]RateLimit{},
		[]TransferEpoch{},
		[]exported.GeneralMessage{},
		0,
	)
}

// Validate checks if the genesis state is valid
func (m GenesisState) Validate() error {
	if err := m.Params.Validate(); err != nil {
		return getValidateError(err)
	}

	for _, chain := range m.Chains {
		if err := chain.Validate(); err != nil {
			return getValidateError(err)
		}
	}

	for _, chainState := range m.ChainStates {
		if err := chainState.Validate(); err != nil {
			return getValidateError(err)
		}
	}

	for _, linkedAddresses := range m.LinkedAddresses {
		if err := linkedAddresses.Validate(); err != nil {
			return getValidateError(err)
		}
	}

	for _, transfer := range m.Transfers {
		if err := transfer.Validate(); err != nil {
			return getValidateError(err)
		}
	}

	if err := m.Fee.Coins.Validate(); err != nil {
		return getValidateError(err)
	}

	for _, feeInfo := range m.FeeInfos {
		if err := feeInfo.Validate(); err != nil {
			return getValidateError(err)
		}
	}

	for _, rateLimit := range m.RateLimits {
		if err := rateLimit.ValidateBasic(); err != nil {
			return getValidateError(err)
		}
	}

	for _, transferEpoch := range m.TransferEpochs {
		if err := transferEpoch.ValidateBasic(); err != nil {
			return getValidateError(err)
		}
	}

	for _, m := range m.Messages {
		if err := m.ValidateBasic(); err != nil {
			return getValidateError(err)
		}
	}

	return nil
}

// GetGenesisStateFromAppState returns x/nexus GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

func getValidateError(err error) error {
	return sdkerrors.Wrapf(err, "genesis state for module %s is invalid", ModuleName)
}
