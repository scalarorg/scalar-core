package types

import (
	utils "github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/btc/exported"
)

const (
	DefaultVaultTag     = "SCALAR"
	DefaultVaultVersion = 0
)

func DefaultGenesisState() GenesisState {
	tag := TagFromAscii(DefaultVaultTag)
	version := VersionFromInt(DefaultVaultVersion)
	return GenesisState{
		Chains: []GenesisState_Chain{
			{
				Params: Params{
					ChainId:             TestnetChainId,
					Chain:               exported.Bitcoin.Name,
					ConfirmationHeight:  2,
					NetworkKind:         Testnet,
					RevoteLockingPeriod: 50,
					VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
					MinVoterCount:       1,
					VotingGracePeriod:   50,
					EndBlockerLimit:     50,
					TransferLimit:       1000,
					VaultTag:            &tag,
					VaultVersion:        &version,
				},
				CommandQueue:        utils.QueueState{},
				ConfirmedStakingTxs: []StakingTx{},
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
