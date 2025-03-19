package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func (s msgServer) ReserveRedeemUtxo(c context.Context, req *types.ReserveRedeemUtxoRequest) (*types.ReserveRedeemUtxoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	//Validate request
	chain, ok := s.nexus.GetChain(ctx, nexus.ChainName(req.Chain))
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if !chainsTypes.IsEvmChain(chain.Name) {
		return nil, fmt.Errorf("chain %s is not a EVM chain", chain.Name)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	reservedTx, err := s.reserveUtxos(ctx, req.Symbol, req.ReqId, req.Amount)
	if err != nil {
		return nil, err
	}

	// Create redeem payload for evm tx
	payload, err := s.createRedeemPayload(ctx, req.Chain.String(), req.Address, req.Symbol, req.Amount, reservedTx)
	if err != nil {
		return nil, err
	}
	// Start signing session for reserve redeem utxos
	keyID, ok := s.multisig.GetCurrentKeyID(ctx, nexus.ChainName(req.Chain))
	if !ok {
		return nil, fmt.Errorf("could not find key ID for '%s'", req.Chain)
	}
	if err := s.multisig.Sign(
		ctx,
		keyID,
		payload,
		types.ModuleName,
		chainsTypes.NewSigMetadata(chainsTypes.SigTx, chain.Name, []byte(req.ReqId)),
	); err != nil {
		return nil, err
	}

	logger := s.Logger(ctx)
	logger.Info("ReserveRedeemUtxoStarted", "reqId", req.ReqId, "amount", req.Amount, "chain", chain.Name, "sender", req.Sender)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			chainsTypes.EventTypeSign,
			sdk.NewAttribute(sdk.AttributeKeyAction, chainsTypes.AttributeValueStart),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, string(req.Amount)),
			sdk.NewAttribute("chain", chain.Name.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, req.Sender.String()),
		),
	)

	return &types.ReserveRedeemUtxoResponse{}, nil
}

func (s msgServer) ConfirmSwitchedPhase(ctx context.Context, req *types.ConfirmSwitchedPhaseRequest) (*types.ConfirmSwitchedPhaseResponse, error) {
	return &types.ConfirmSwitchedPhaseResponse{}, nil
}
