package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gethParams "github.com/ethereum/go-ethereum/params"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	utils "github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/chains/exported"
)

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Chains: []GenesisState_Chain{
			{
				Params:              DefaultChainParams(BTCMainnetChainID, exported.Bitcoin.Name, vault.NetworkKindMainnet, map[string]string{}),
				CommandQueue:        utils.QueueState{},
				ConfirmedSourceTxs:  []SourceTx{},
				CommandBatches:      []CommandBatchMetadata{},
				Events:              []Event{},
				Tokens:              []ERC20TokenMetadata{},
				ConfirmedEventQueue: utils.QueueState{},
			},
			{
				Params:              DefaultChainParams(sdk.NewIntFromBigInt(gethParams.MainnetChainConfig.ChainID), exported.Ethereum.Name, vault.NetworkKindMainnet, map[string]string{}),
				CommandQueue:        utils.QueueState{},
				ConfirmedSourceTxs:  []SourceTx{},
				CommandBatches:      []CommandBatchMetadata{},
				Events:              []Event{},
				Tokens:              []ERC20TokenMetadata{},
				ConfirmedEventQueue: utils.QueueState{},
			},
		},
	}
}

func (data GenesisState) Validate() error {
	for _, chain := range data.Chains {
		if err := chain.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func NewGenesisState(chains []GenesisState_Chain) *GenesisState {
	return &GenesisState{
		Chains: chains,
	}
}

func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
