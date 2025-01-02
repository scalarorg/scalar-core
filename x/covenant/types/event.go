package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewSigningStarted is the constructor for event signing started
func NewCreateAndSigningPsbtStarted(sigID uint64, key multisigTypes.Key, batchPsbtPayload []PsbtPayload, requestingModule string, chainName nexus.ChainName) *CreateAndSigningPsbtStarted {
	return &CreateAndSigningPsbtStarted{
		Module:           ModuleName,
		Chain:            chainName,
		SigID:            sigID,
		KeyID:            key.GetID(),
		PubKeys:          key.GetPubKeys(),
		BatchPsbtPayload: batchPsbtPayload,
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
