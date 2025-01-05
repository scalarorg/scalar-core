package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewSigningPsbtStarted is the constructor for event signing started
func NewSigningPsbtStarted(sigID uint64, key multisigTypes.Key, psbt []byte, requestingModule string, chainName nexus.ChainName) *SigningPsbtStarted {
	return &SigningPsbtStarted{
		Module:           ModuleName,
		Chain:            chainName,
		SigID:            sigID,
		KeyID:            key.GetID(),
		PubKeys:          key.GetPubKeys(),
		Psbt:             psbt,
		RequestingModule: requestingModule,
	}
}

// NewSigningPsbtExpired is the constructor for event signing expired
func NewSigningPsbtExpired(sigID uint64) *SigningPsbtExpired {
	return &SigningPsbtExpired{
		Module: ModuleName,
		SigID:  sigID,
	}
}

// NewSigningCompleted is the constructor for event signing completed
func NewSigningPsbtCompleted(sigID uint64) *SigningPsbtCompleted {
	return &SigningPsbtCompleted{
		Module: ModuleName,
		SigID:  sigID,
	}
}

// NewTapscriptSigsSubmitted is the constructor for event tapscript sig submitted
func NewTapscriptSigsSubmitted(sigID uint64, participant sdk.ValAddress, tapscriptSigs *exported.TapScriptSigList) *TapScriptSigsSubmitted {
	return &TapScriptSigsSubmitted{
		Module:        ModuleName,
		SigID:         sigID,
		Participant:   participant,
		TapScriptSigs: tapscriptSigs,
	}
}
