package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func (s msgServer) RegisterCustodianGroup(c context.Context, req *types.RegisterCustodianGroupRequest) (*types.RegisterCustodianGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, err := s.validateEvmChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	_, ok := s.covenant.GetCustodianGroup(ctx, req.CustodianGroupUID)
	if !ok {
		return nil, fmt.Errorf("custodian group %s not found", req.CustodianGroupUID)
	}
	keyID, ok := s.multisig.GetCurrentKeyID(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("current key not set for chain %s", req.Chain)
	}
	keeper, err := s.ForChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}
	chainID := funcs.MustOk(keeper.GetChainID(ctx))

	cmd := types.NewRegisterCustodianGroupCommand(
		chainID,
		keyID,
		req.CustodianGroupUID,
	)
	if err := keeper.EnqueueCommand(ctx, cmd); err != nil {
		return nil, err
	}
	return &types.RegisterCustodianGroupResponse{}, nil
}

func (s msgServer) validateEvmChain(ctx sdk.Context, chain nexus.ChainName) (*nexus.Chain, error) {
	if !types.IsEvmChain(chain) {
		return nil, fmt.Errorf("chain %s is not a EVM chain", chain)
	}
	c, ok := s.nexus.GetChain(ctx, chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", chain)
	}
	if err := validateChainActivated(ctx, s.nexus, c); err != nil {
		return nil, err
	}

	return &c, nil
}
