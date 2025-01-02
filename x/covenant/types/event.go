package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
)

// NewSigningStarted is the constructor for event signing started
func NewSigningStarted(sigID uint64, key multisigTypes.Key, psbt Psbt, requestingModule string) *SigningStarted {
	return &SigningStarted{
		Module:           ModuleName,
		SigID:            sigID,
		KeyID:            key.GetID(),
		PubKeys:          key.GetPubKeys(),
		Psbt:             psbt,
		RequestingModule: requestingModule,
	}
}

// NewSigningExpired is the constructor for event signing expired
func NewSigningExpired(sigID uint64) *SigningExpired {
	return &SigningExpired{
		Module: ModuleName,
		SigID:  sigID,
	}
}

// NewSigningCompleted is the constructor for event signing completed
func NewSigningCompleted(sigID uint64) *SigningCompleted {
	return &SigningCompleted{
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
