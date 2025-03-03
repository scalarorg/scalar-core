package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	covenanttypes "github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Querier)(nil)

type Querier struct {
	keeper   *Keeper
	covenant types.CovenantKeeper
}

func NewGRPCQuerier(keeper *Keeper, covenant types.CovenantKeeper) *Querier {
	return &Querier{keeper: keeper, covenant: covenant}
}

// GovernanceKey returns the xmultisig governance key
func (k *Querier) Protocols(c context.Context, req *types.ProtocolsRequest) (*types.ProtocolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	protocols, ok := k.keeper.findProtocols(ctx, req)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "protocols not found")
	}

	protocolDetails := make([]*types.ProtocolDetails, len(protocols))
	for i, protocol := range protocols {
		custodianGr, ok := k.covenant.GetCustodianGroup(ctx, protocol.CustodianGroupUID)
		if !ok {
			ctx.Logger().Error("custodian group not found", "protocol", protocol.Asset.Name, "custodian group uid", protocol.CustodianGroupUID)
			return nil, status.Errorf(codes.NotFound, "custodian group not found")
		}
		protocolDetails[i] = mapProtocolToProtocolDetails(protocol, custodianGr)
	}

	return &types.ProtocolsResponse{
		Protocols: protocolDetails,
		Total:     uint64(len(protocols)),
	}, nil
}

func (q *Querier) Protocol(c context.Context, req *types.ProtocolRequest) (*types.ProtocolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := req.ValidateBasic()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	var protocol *types.Protocol
	if req.Symbol != "" {
		protocol, err = q.keeper.FindProtocolByExternalSymbol(ctx, req.Symbol)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "protocol not found")
		}

		custodianGr, ok := q.covenant.GetCustodianGroup(ctx, protocol.CustodianGroupUID)
		if !ok {
			return nil, status.Errorf(codes.NotFound, "custodian group not found")
		}

		return &types.ProtocolResponse{
			Protocol: mapProtocolToProtocolDetails(protocol, custodianGr),
		}, nil
	}

	if req.Address != "" {
		protocol, err = q.keeper.FindProtocolByInternalAddress(ctx, req.OriginChain, req.MinorChain, req.Address)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "protocol not found")
		}

		custodianGr, ok := q.covenant.GetCustodianGroup(ctx, protocol.CustodianGroupUID)
		if !ok {
			return nil, status.Errorf(codes.NotFound, "custodian group not found")
		}

		return &types.ProtocolResponse{
			Protocol: mapProtocolToProtocolDetails(protocol, custodianGr),
		}, nil
	}

	// This should never happen because of the validation above, but it enstures in case of the validation is not working
	return nil, status.Errorf(codes.NotFound, "protocol not found")
}

func mapProtocolToProtocolDetails(protocol *types.Protocol, custodianGr *covenanttypes.CustodianGroup) *types.ProtocolDetails {
	return &types.ProtocolDetails{
		BitcoinPubkey:     protocol.BitcoinPubkey,
		ScalarPubkey:      protocol.ScalarPubkey,
		ScalarAddress:     protocol.ScalarAddress,
		Name:              protocol.Name,
		Tag:               protocol.Tag,
		Attributes:        protocol.Attributes,
		Status:            protocol.Status,
		CustodianGroupUID: protocol.CustodianGroupUID,
		Asset:             protocol.Asset,
		Chains:            protocol.Chains,
		Avatar:            protocol.Avatar,
		CustodianGroup:    custodianGr,
	}
}
