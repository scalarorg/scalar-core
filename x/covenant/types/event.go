package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisigExported "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

const (
	AttributeKeyChain    = "chain"
	AttributeKeyDataHash = "dataHash"
	AttributeCommandId   = "commandId"
)

const (
	AttributeValueStart   = "start"
	AttributeValueConfirm = "confirm"
)

// NewSigningPsbtStarted is the constructor for event signing started
func NewSigningPsbtStarted(sigID uint64, key multisigTypes.Key, multiPsbt []exported.Psbt, requestingModule string, chainName nexus.ChainName) *SigningPsbtStarted {
	return &SigningPsbtStarted{
		Module:           ModuleName,
		Chain:            chainName,
		SigID:            sigID,
		KeyID:            key.GetID(),
		PubKeys:          key.GetPubKeys(),
		MultiPsbt:        multiPsbt,
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
func NewTapscriptSigsSubmitted(sigID uint64, participant sdk.ValAddress, list []*exported.TapScriptSigsMap) *TapScriptSigsSubmitted {
	return &TapScriptSigsSubmitted{
		Module:                 ModuleName,
		SigID:                  sigID,
		Participant:            participant,
		ListOfTapScriptSigsMap: list,
	}
}

// NewKeyRotated is the constructor for event key rotated
func NewKeyRotated(chain nexus.ChainName, keyID multisigExported.KeyID) *KeyRotated {
	return &KeyRotated{
		Module: ModuleName,
		Chain:  chain,
		KeyID:  keyID,
	}
}

func (m Event) ValidateBasic() error {
	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid source chain")
	}
	// TODO: validate event type

	return nil
}

func (m VoteEvents) ValidateBasic() error {
	if err := m.Chain.Validate(); err != nil {
		return err
	}

	for _, event := range m.Events {
		if err := event.ValidateBasic(); err != nil {
			return err
		}

		if event.Chain != m.Chain {
			return fmt.Errorf("events are not from the same source chain")
		}
	}

	return nil
}

// NewRedeemTxsConfirmed is the constructor for event redeem txs confirmed
func (e *RedeemTxsConfirmed) ValidateBasic() error {
	// TODO: validate
	return nil
}
func (e *SwitchedPhaseConfirmed) ValidateBasic() error {
	// TODO: validate
	if e.FromPhase == e.ToPhase {
		return fmt.Errorf("from phase and to phase are the same")
	}
	if e.CustodianGroupUID.IsZero() {
		return fmt.Errorf("custodian group uid is zero")
	}
	return nil
}

func (e *IntializeUtxoSnapshotCompleted) ValidateBasic() error {
	// TODO: validate
	return nil
}
