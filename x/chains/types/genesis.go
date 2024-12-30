package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gethParams "github.com/ethereum/go-ethereum/params"
	utils "github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func DefaultChainParams(chainId sdk.Int, chain nexus.ChainName, networkKind NetworkKind, metadata map[string]string) Params {
	return Params{
		ChainID:             chainId,
		Chain:               chain,
		ConfirmationHeight:  2,
		NetworkKind:         networkKind,
		RevoteLockingPeriod: 50,
		VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
		MinVoterCount:       1,
		CommandsGasLimit:    5000000,
		VotingGracePeriod:   50,
		EndBlockerLimit:     50,
		TransferLimit:       1000,
		ProcessingTxsWindowSize: 7,
		Metadata:               metadata,
	}
}
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Chains: []GenesisState_Chain{
			{
				Params:              DefaultChainParams(BTCMainnetChainID, exported.Bitcoin.Name, Mainnet, map[string]string{}),
				CommandQueue:        utils.QueueState{},
				ConfirmedSourceTxs:  []SourceTx{},
				CommandBatches:      []CommandBatchMetadata{},
				Events:              []Event{},
				ConfirmedEventQueue: utils.QueueState{},
			},
			{
				Params:              DefaultChainParams(sdk.NewIntFromBigInt(gethParams.MainnetChainConfig.ChainID), exported.Ethereum.Name, Mainnet, map[string]string{}),
				CommandQueue:        utils.QueueState{},
				ConfirmedSourceTxs:  []SourceTx{},
				CommandBatches:      []CommandBatchMetadata{},
				Events:              []Event{},
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
