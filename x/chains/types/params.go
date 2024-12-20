package types

import (
	fmt "fmt"

	utils "github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/log"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyChainName           = []byte("chainName")
	KeyConfirmationHeight  = []byte("confirmationHeight")
	KeyNetworkKind         = []byte("networkKind")
	KeyRevoteLockingPeriod = []byte("revoteLockingPeriod")
	KeyChainID             = []byte("chainId")
	KeyVotingThreshold     = []byte("votingThreshold")
	KeyMinVoterCount       = []byte("minVoterCount")
	KeyVotingGracePeriod   = []byte("votingGracePeriod")
	KeyEndBlockerLimit     = []byte("endBlockerLimit")
	KeyTransferLimit       = []byte("transferLimit")
)

func KeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func (Params) Validate() error {
	// TODO: Implement validation
	log.Debug("Not implemented params validation")
	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyChainName, &m.Chain, validateChainName),
		params.NewParamSetPair(KeyConfirmationHeight, &m.ConfirmationHeight, validateConfirmationHeight),
		params.NewParamSetPair(KeyNetworkKind, &m.NetworkKind, validateNetworkKind),
		params.NewParamSetPair(KeyRevoteLockingPeriod, &m.RevoteLockingPeriod, validateRevoteLockingPeriod),
		params.NewParamSetPair(KeyChainID, &m.ChainID, validateChainId),
		params.NewParamSetPair(KeyVotingThreshold, &m.VotingThreshold, validateVotingThreshold),
		params.NewParamSetPair(KeyMinVoterCount, &m.MinVoterCount, validateMinVoterCount),
		params.NewParamSetPair(KeyVotingGracePeriod, &m.VotingGracePeriod, validateVotingGracePeriod),
		params.NewParamSetPair(KeyEndBlockerLimit, &m.EndBlockerLimit, validateEndBlockerLimit),
		params.NewParamSetPair(KeyTransferLimit, &m.TransferLimit, validateTransferLimit),
	}
}

func validateChainName(i interface{}) error {
	chainName, ok := i.(nexus.ChainName)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return chainName.Validate()
}

func validateConfirmationHeight(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNetworkKind(i interface{}) error {
	networkKind, ok := i.(NetworkKind)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return networkKind.Validate()
}

func validateRevoteLockingPeriod(i interface{}) error {
	period, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if period < 1 {
		return fmt.Errorf("revote locking period must be greater than 0")
	}
	return nil
}

func validateChainId(i interface{}) error {
	chainId, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !chainId.IsPositive() {
		return fmt.Errorf("chain id must be positive")
	}
	return nil
}

func validateVotingThreshold(i interface{}) error {
	threshold, ok := i.(utils.Threshold)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return threshold.Validate()
}

func validateMinVoterCount(i interface{}) error {
	count, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if count < 1 {
		return fmt.Errorf("min voter count must be greater than 0")
	}
	return nil
}

func validateVotingGracePeriod(i interface{}) error {
	period, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if period < 1 {
		return fmt.Errorf("voting grace period must be greater than 0")
	}
	return nil
}

func validateEndBlockerLimit(i interface{}) error {
	limit, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if limit < 1 {
		return fmt.Errorf("end blocker limit must be greater than 0")
	}
	return nil
}

func validateTransferLimit(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
