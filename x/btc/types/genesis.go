package types

import (
	utils "github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

const (
	DefaultVaultTag     = "SCALAR"
	DefaultVaultVersion = 0
)

func DefaultGenesisState() *GenesisState {
	tag := TagFromAscii(DefaultVaultTag)
	version := VersionFromInt(DefaultVaultVersion)
	return &GenesisState{
		Chains: []GenesisState_Chain{
			{
				Params: Params{
					ChainId:             TestnetChainId,
					ChainName:           exported.ChainName("bitcoin-testnet4"),
					ConfirmationHeight:  2,
					NetworkKind:         Testnet,
					RevoteLockingPeriod: 50,
					VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
					MinVoterCount:       1,
					VotingGracePeriod:   50,
					EndBlockerLimit:     50,
					TransferLimit:       1000,
				},
			},
		},
		VaultTag:     &tag,
		VaultVersion: &version,
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

func NewGenesisState(params Params, tag VaultTag, version VaultVersion) *GenesisState {
	return &GenesisState{
		Chains: []GenesisState_Chain{
			{
				Params: params,
			},
		},
		VaultTag:     &tag,
		VaultVersion: &version,
	}
}
