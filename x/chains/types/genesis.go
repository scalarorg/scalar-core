package types

import (
	gethParams "github.com/ethereum/go-ethereum/params"
	utils "github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/chains/exported"
)

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Chains: []GenesisState_Chain{
			{
				Params: Params{
					ChainID:             BTCMainnetChainID,
					Chain:               exported.Bitcoin.Name,
					ConfirmationHeight:  2,
					NetworkKind:         Mainnet,
					RevoteLockingPeriod: 50,
					VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
					MinVoterCount:       1,
					VotingGracePeriod:   50,
					EndBlockerLimit:     50,
					TransferLimit:       1000,
				},
				CommandQueue:        utils.QueueState{},
				ConfirmedStakingTxs: []StakingTx{},
				CommandBatches:      []CommandBatchMetadata{},
				Events:              []Event{},
				ConfirmedEventQueue: utils.QueueState{},
			},
			{
				Params: Params{
					Chain:               exported.Ethereum.Name,
					ChainID:             gethParams.MainnetChainConfig.ChainID.Uint64(),
					RevoteLockingPeriod: 50,
					VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
					VotingGracePeriod:   3,
					MinVoterCount:       1,
					EndBlockerLimit:     50,
					TransferLimit:       50,
				},
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