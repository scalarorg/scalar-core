package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func NewCreatingPsbtStarted() *CreatingPsbtStarted {
	return &CreatingPsbtStarted{}
}

// NewSigningPsbtStarted is the constructor for event signing started
func NewSigningPsbtStarted(sigID uint64, key multisigTypes.Key, batchPsbtPayload []PsbtPayload, requestingModule string, chainName nexus.ChainName) *SigningPsbtStarted {
	return &SigningPsbtStarted{
		Module:           ModuleName,
		Chain:            chainName,
		SigID:            sigID,
		KeyID:            key.GetID(),
		PubKeys:          key.GetPubKeys(),
		BatchPsbtPayload: batchPsbtPayload,
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

// NewTapscriptSigSubmitted is the constructor for event tapscript sig submitted
func NewTapscriptSigSubmitted(sigID uint64, participant sdk.ValAddress, tapscriptSig *exported.TapScriptSig) *TapScriptSigSubmitted {
	return &TapScriptSigSubmitted{
		Module:       ModuleName,
		SigID:        sigID,
		Participant:  participant,
		TapScriptSig: tapscriptSig,
	}
}
